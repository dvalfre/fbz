package http

import (
	"encoding/json"
	"fmt"

	"github.com/ess/fbz/pkg/fbz"
)

type CaseService struct {
	driver fbz.Driver
}

func NewCaseService(driver fbz.Driver) *CaseService {
	return &CaseService{driver: driver}
}

func (s *CaseService) All(query string) []*fbz.Case {
	cols := []string{"sTitle", "sProject", "sArea", "sStatus", "dblStoryPts", "sPriority"}

	cmd, err := json.Marshal(
		&searchCmd{
			Cmd:   "search",
			Q:     query,
			Cols:  cols,
			Token: s.driver.Token(),
		},
	)

	if err == nil {
		r := s.driver.Post("/f/api/0/jsonapi", nil, []byte(cmd))
		if r.Okay() {
			wrapper := caseWrapper{}
			jErr := json.Unmarshal(r.Data, &wrapper)
			if jErr == nil {
				return wrapper.Data.Cases
			}
		}
	}

	return make([]*fbz.Case, 0)
}

func (s *CaseService) Get(caseID int) (*fbz.Case, error) {
	cols := []string{"sTitle", "sStatus", "events"}

	cmd, err := json.Marshal(
		&getCmd{
			Cmd:   "search",
			Q:     caseID,
			Cols:  cols,
			Token: s.driver.Token(),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("could not build api command")
	}

	r := s.driver.Post("/f/api/0/jsonapi", nil, []byte(cmd))
	if !r.Okay() {
		return nil, fmt.Errorf("ticket not found")
	}

	wrapper := caseWrapper{}

	err = json.Unmarshal(r.Data, &wrapper)
	if err != nil {
		return nil, fmt.Errorf("could not decode server response")
	}

	collection := wrapper.Data.Cases
	if len(collection) == 0 {
		return nil, fmt.Errorf("ticket not found")
	}

	return collection[0], nil
}

func (s *CaseService) Update(caseID int, message string) (*fbz.Case, error) {
	c, err := s.Get(caseID)
	if err != nil {
		return nil, err
	}

	cmd, err := json.Marshal(
		&updateCmd{
			Cmd:     "edit",
			CaseID:  c.ID,
			Message: message,
			Token:   s.driver.Token(),
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

type searchCmd struct {
	Cmd   string   `json:"cmd"`
	Q     string   `json:"q"`
	Cols  []string `json:"cols"`
	Token string   `json:"token"`
}

type getCmd struct {
	Cmd   string   `json:"cmd"`
	Q     int      `json:"q"`
	Cols  []string `json:"cols"`
	Token string   `json:"token"`
}

type updateCmd struct {
	Cmd     string `json:"cmd"`
	CaseID  int    `json:"ixBug"`
	Message string `json:"sEvent"`
	Token   string `json:"token"`
}

type caseWrapper struct {
	Data *caseDetailsData `json:"data"`
}

type caseDetailsData struct {
	Cases []*fbz.Case `json:"cases"`
}
