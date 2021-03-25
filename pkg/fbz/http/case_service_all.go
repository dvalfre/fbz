package http

import (
	"encoding/json"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) All(query string) []*fbz.Case {
	cmd, err := json.Marshal(
		&searchCmd{
			Cmd:   "search",
			Q:     query,
			Cols:  caseCols,
			Token: s.driver.Token(),
		},
	)

	if err == nil {
		r := s.driver.Post("/f/api/0/jsonapi", []byte(cmd))
		if r.Okay() {
			wrapper := caseWrapper{}
			jErr := json.Unmarshal(r.Data(), &wrapper)
			if jErr == nil {
				return wrapper.Data.Cases
			}
		}
	}

	return make([]*fbz.Case, 0)
}

type searchCmd struct {
	Cmd   string   `json:"cmd"`
	Q     string   `json:"q"`
	Cols  []string `json:"cols"`
	Token string   `json:"token"`
}
