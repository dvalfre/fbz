package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Accept(caseID int, message string) (*fbz.Case, error) {
	c, err := s.Get(caseID)
	if err != nil {
		return nil, err
	}

	cmd, err := json.Marshal(
		&acceptCmd{
			Cmd:     "close",
			CaseID:  c.ID,
			Message: message,
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

type acceptCmd struct {
	Cmd     string `json:"cmd"`
	CaseID  int    `json:"ixBug"`
	Message string `json:"sEvent"`
	Token   string `json:"token"`
}
