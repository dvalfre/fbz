package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Estimate(caseID int, points int) (*fbz.Case, error) {
	c, err := s.Get(caseID)
	if err != nil {
		return nil, err
	}

	cmd, err := json.Marshal(
		&estimateCmd{
			Cmd:    "edit",
			CaseID: c.ID,
			Points: points,
			Token:  s.driver.Token(),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("could not build api command")
	}

	r := s.driver.Post("/f/api/0/jsonapi", nil, []byte(cmd))
	if !r.Okay() {
		return nil, fmt.Errorf("the api reported an error")
	}

	return s.Get(caseID)
}

type estimateCmd struct {
	Cmd    string `json:"cmd"`
	CaseID int    `json:"ixBug"`
	Points int    `json:"dblStoryPts"`
	Token  string `json:"token"`
}
