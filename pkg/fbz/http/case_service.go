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

var caseCols = []string{
	"sTitle",
	"sProject",
	"sArea",
	"sStatus",
	"dblStoryPts",
	"sPriority",
	"sCategory",
	"events",
}

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

	fmt.Println("DEBUG:", string(cmd))

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

type createCmd struct {
	Cmd      string `json:"cmd"`
	Title    string `json:"sTitle"`
	Message  string `json:"sEvent"`
	Category string `json:"sCategory"`
	Project  string `json:"sProject"`
	Area     string `json:"sArea"`
	Token    string `json:"token"`
}

type resolveCmd struct {
	Cmd     string `json:"cmd"`
	CaseID  int    `json:"ixBug"`
	Message string `json:"sEvent"`
	Status  string `json:"sStatus"`
	Token   string `json:"token"`
}

type caseWrapper struct {
	Data *caseDetailsData `json:"data"`
}

type singleCaseWrapper struct {
	Data *singleCaseDetailsData `json:"data"`
}

type caseDetailsData struct {
	Cases []*fbz.Case `json:"cases"`
}

type singleCaseDetailsData struct {
	Case *fbz.Case `json:"case"`
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
