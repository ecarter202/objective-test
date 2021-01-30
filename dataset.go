package main

type DataSet struct {
	Applicants []*Applicant `json:"applicants"`
	Jobs       []*Job       `json:"jobs"`
	Skills     []*Skill     `json:"skills"`
}
