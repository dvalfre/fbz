package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Assign(caseID int, person string) (*fbz.Case, error) {
	c, err := s.Get(caseID)
	if err != nil {
		return nil, err
	}

	cmd, err := json.Marshal(
		&assignCmd{
			Cmd:    "assign",
			CaseID: c.ID,
			Person: person,
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

type assignCmd struct {
	Cmd    string `json:"cmd"`
	CaseID int    `json:"ixBug"`
	Person string `json:"sPersonAssignedTo"`
	Token  string `json:"token"`
}
