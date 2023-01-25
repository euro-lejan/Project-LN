package models

import (
	"gorm.io/gorm"
	"time"
)

type Saveunit struct {
	ID   *uint      `gorm:"primary_key; auto_increment; index;" json:"save_uintId"`
	Unit *float32   `json:"unit"`
	Date *string`json:"date"`

	CreatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	UpdatedAt *time.Time     `gorm:"DEFAULT:now()" json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
