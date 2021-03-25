package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Get(caseID int) (*fbz.Case, error) {
	cmd, err := json.Marshal(
		&getCmd{
			Cmd:   "search",
			Q:     caseID,
			Cols:  caseCols,
			Token: s.driver.Token(),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("could not build api command")
	}

	r := s.driver.Post("/f/api/0/jsonapi", []byte(cmd))
	if !r.Okay() {
		return nil, fmt.Errorf("ticket not found")
	}

	wrapper := caseWrapper{}

	err = json.Unmarshal(r.Data(), &wrapper)
	if err != nil {
		return nil, fmt.Errorf("could not decode server response")
	}

	collection := wrapper.Data.Cases
	if len(collection) == 0 {
		return nil, fmt.Errorf("ticket not found")
	}

	return collection[0], nil
}

type getCmd struct {
	Cmd   string   `json:"cmd"`
	Q     int      `json:"q"`
	Cols  []string `json:"cols"`
	Token string   `json:"token"`
}
