package fbz

type Project struct {
	ID   int    `json:"ixProject"`
	Name string `json:"sProject"`

	Inbox   bool `json:"fInbox"`
	Deleted bool `json:"fDeleted"`

	OwnerID    int    `json:"ixPerson"`
	Owner      string `json:"sPerson"`
	OwnerEmail string `json:"sEmail"`
	OwnerPhone string `json:"sPhone"`

	WorkflowID int `json:"ixWorkflow"`
}

type ProjectService interface {
	All() []*Project
	ByName(string) (*Project, error)
}
