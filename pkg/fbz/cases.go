package fbz

type Case struct {
	ID          int      `json:"ixBug"`
	Title       string   `json:"sTitle"`
	Status      string   `json:"sStatus"`
	Priority    string   `json:"sPriority"`
	Points      int      `json:"dblStoryPts"`
	ProjectName string   `json:"sProject"`
	AreaName    string   `json:"sArea"`
	Events      []*Event `json:"events"`
}

type Event struct {
	ID           int    `json:"ixBugEvent"`
	CaseID       int    `json:"ixBug"`
	Code         int    `json:"evt"`
	Verb         string `json:"sVerb"`
	PersonID     int    `json:"ixPerson"`
	AssignedToID int    `json:"ixPersonAssignedTo"`
	CreatedAt    string `json:"dt"`
	Text         string `json:"s"`
	Changes      string `json:"sChanges"`
	Format       string `json:"sFormat"`
	Description  string `json:"evtDescription"`
	Creator      string `json:"sPerson"`
	HTML         string `json:"sHtml"`
}

type CaseService interface {
	All(map[string]string) []*Case
	Get(int) (*Case, error)
	Update(int, string) (*Case, error)
}
