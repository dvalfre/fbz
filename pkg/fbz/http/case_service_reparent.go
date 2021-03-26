package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Reparent(caseID int, parentID int) (*fbz.Case, error) {
	c, err := s.Get(caseID)
	if err != nil {
		return nil, err
	}

	parent, err := s.Get(parentID)
	if err != nil {
		return nil, fmt.Errorf("could not get parent")
	}

	cmd, err := json.Marshal(
		&reparentCmd{
			Cmd:      "edit",
			CaseID:   c.ID,
			ParentID: parent.ID,
			Token:    s.driver.Token(),
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

type reparentCmd struct {
	Cmd      string `json:"cmd"`
	CaseID   int    `json:"ixBug"`
	ParentID int    `json:"ixBugParent"`
	Token    string `json:"token"`
}
