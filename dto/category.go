package dto

type CategoryHierarchy struct {
	ID       string              `json:"id"`
	Name     string              `json:"name"`
	Slug     string              `json:"slug"`
	Icon     string              `json:"icon"`
	Children []CategoryHierarchy `json:"children"`
}
