package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Resolve(caseID int, reject bool, message string) (*fbz.Case, error) {
	c, err := s.Get(caseID)
	if err != nil {
		return nil, err
	}

	cmd, err := json.Marshal(
		&resolveCmd{
			Cmd:     "resolve",
			CaseID:  c.ID,
			Message: message,
			Status:  resolutionText(c.Category, reject),
			Token:   s.driver.Token(),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("could not build api command")
	}

	r := s.driver.Post("/f/api/0/jsonapi", []byte(cmd))
	if !r.Okay() {
		return nil, fmt.Errorf("the api reported an error")
	}

	return s.Get(caseID)
}

type resolveCmd struct {
	Cmd     string `json:"cmd"`
	CaseID  int    `json:"ixBug"`
	Message string `json:"sEvent"`
	Status  string `json:"sStatus"`
	Token   string `json:"token"`
}

func resolutionText(category string, reject bool) string {
	if !reject {
		return "Resolved"
	}

	return rejectionText(category, reject)
}

func rejectionText(category string, reject bool) string {
	if reject {
		switch category {
		case "Task":
			return "Resolved (Won't Implement)"
		case "Bug":
			return "Resolved (Won't Fix)"
		case "Feature":
			return "Resolved (Duplicate)"
		case "Inquiry":
			return "Resolved (Won't Respond)"
		}

		return "Resolved"
	}

	return resolutionText(category, reject)
}
