package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type RESTLink struct {
	Self string `json:"_self"`
}

type BaseModel struct {
	Links     RESTLink   `json:"_links" gorm:"-"`
	UUID      string     `gorm:"primary_key" json:"uuid"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type StorableModel interface {
	BeforeCreate(scope *gorm.Scope) error
}

func (baseModel BaseModel) BeforeCreate(scope *gorm.Scope) error {
	uid, err := uuid.NewUUID()

	if err != nil {
		return err
	}
	scope.SetColumn("UUID", uid.String())
	return nil
}
