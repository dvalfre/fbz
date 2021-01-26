package fbz

type Area struct {
	ID   int    `json:"ixArea"`
	Name string `json:"sArea"`

	ProjectID int `json:"ixProject"`
}

type AreaService interface {
	All(*Project) []*Area
	ByName(*Project, string) (*Area, error)
}
