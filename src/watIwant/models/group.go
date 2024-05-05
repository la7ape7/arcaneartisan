package models

type Group struct {
	BaseModel
	Name  string `json:""`
	Users []User `json:""`
}
