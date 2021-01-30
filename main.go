package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/gocraft/web"
)

const (
	port = 8080
)

type Context struct{}

var (
	applicantsTemplate *template.Template
	dataset            *DataSet
	path               string
	err                error
)

func init() {
	path = os.Args[1]
	extX := strings.Split(path, ".")
	ext := extX[len(extX)-1]
	selectedDataSource = supportedFileTypes[ext]

	if strings.TrimSpace(path) == "" {
		log.Fatal("file path is required")
	} else if selectedDataSource == nil {
		log.Fatal("invalid file type supplied")
	}
}

func main() {
	applicantsTemplate, err = template.ParseFiles("templates/applicants.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := selectedDataSource.Load(path); err != nil {
		log.Fatal(err)
	}

	ctx := Context{}
	router := web.New(ctx)
	router.Get("/", ctx.showApplicants)

	server := http.Server{
		ReadTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 15,
		WriteTimeout:      time.Second * 60,
		IdleTimeout:       time.Second * 30,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
		Handler:           router,
		Addr:              fmt.Sprintf(":%d", port),
	}

	fmt.Printf("Serving HTML on port :%d\n", port)

	log.Fatal(server.ListenAndServe())
}

func (ctx Context) showApplicants(rw web.ResponseWriter, req *web.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("content-type", "text/html; charset=UTF-8")

	formattedDataSet := &FormattedDataSet{
		Jobs: make(map[int]*FormattedJob),
	}

	var seenSkills = make(map[string]bool)
	for _, job := range dataset.Jobs {
		formattedDataSet.Jobs[job.ID] = &FormattedJob{
			ID:   job.ID,
			Name: job.Name,
		}
	}

	for _, app := range dataset.Applicants {
		formattedApp := new(FormattedApplicant)
		formattedApp.Applicant = app
		formattedDataSet.Jobs[app.JobID].Applicants = append(formattedDataSet.Jobs[app.JobID].Applicants, formattedApp)
		for _, skill := range dataset.Skills {
			if !seenSkills[skill.Name] {
				formattedDataSet.UniqueSkillTotal++
				seenSkills[skill.Name] = true
			}
			if skill.ApplicantID == app.ID {
				formattedApp.Skills = append(formattedApp.Skills, skill)
				formattedDataSet.Jobs[app.JobID].TotalSkillCount++
			}
		}
		formattedDataSet.ApplicantTotal++
	}

	applicantsTemplate.Execute(rw, formattedDataSet)
}
