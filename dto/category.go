package dto

type CategoryNode struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Slug     string         `json:"slug"`
	Icon     string         `json:"icon"`
	Children []CategoryNode `json:"children,omitempty"`
}
