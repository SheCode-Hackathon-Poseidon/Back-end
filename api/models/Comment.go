package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
    gorm.Model
    ProgressID int
    UserID     int
    Content    string `gorm:"type:varchar"`
}