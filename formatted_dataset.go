package main

type FormattedDataSet struct {
	Jobs             map[int]*FormattedJob `json:"jobs"`
	ApplicantTotal   int                   `json:"applicant_total"`
	UniqueSkillTotal int                   `json:"unique_skill_total"`
}

type FormattedJob struct {
	ID              int                   `json:"id"`
	Name            string                `json:"name"`
	Applicants      []*FormattedApplicant `json:"applicants"`
	TotalSkillCount int                   `json:"total_skill_count"`
}

type FormattedApplicant struct {
	*Applicant
	Skills []*Skill `json:"skills"`
}
