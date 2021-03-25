package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Start(caseID int) (*fbz.Case, error) {
	c, err := s.Get(caseID)
	if err != nil {
		return nil, err
	}

	cmd, err := json.Marshal(
		&startCmd{
			Cmd:    "edit",
			CaseID: c.ID,
			Status: "In Progress",
			Token:  s.driver.Token(),
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

type startCmd struct {
	Cmd    string `json:"cmd"`
	CaseID int    `json:"ixBug"`
	Status string `json:"sStatus"`
	Token  string `json:"token"`
}
