package main

type Applicant struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Website     string `json:"website"`
	CoverLetter string `json:"cover_letter"`
	JobID       int    `json:"job_id"`
}
