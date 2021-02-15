package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

func (s *CaseService) Create(projectName string, areaName string, title string, category string, message string) (*fbz.Case, error) {
	cmd, err := json.Marshal(
		&createCmd{
			Cmd:      "new",
			Project:  projectName,
			Area:     areaName,
			Category: category,
			Title:    title,
			Message:  message,
			Token:    s.driver.Token(),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("could not build api command")
	}

	r := s.driver.Post("/f/api/0/jsonapi", nil, []byte(cmd))
	if !r.Okay() {
		return nil, fmt.Errorf("the api reported an error")
	}

	wrapper := &singleCaseWrapper{}

	err = json.Unmarshal(r.Data, &wrapper)
	if err != nil {
		return nil, fmt.Errorf("could not decode api response")
	}

	return s.Get(wrapper.Data.Case.ID)
}

type createCmd struct {
	Cmd      string `json:"cmd"`
	Title    string `json:"sTitle"`
	Message  string `json:"sEvent"`
	Category string `json:"sCategory"`
	Project  string `json:"sProject"`
	Area     string `json:"sArea"`
	Token    string `json:"token"`
}
