package models

type Item struct {
	BaseModel
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
}
