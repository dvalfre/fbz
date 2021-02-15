package http

import (
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
	"sPersonAssignedTo",
	"events",
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
