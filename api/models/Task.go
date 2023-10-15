package models

import (
	"github.com/jinzhu/gorm"

)

type Task struct {
    gorm.Model
    ProjectID int `json:"project_id"`
    Title     string `gorm:"type:varchar" json: "title"`
    Deadline  string `json: "deadline"`
}

func (p *Task) Create(db *gorm.DB) (*Task, error) {
	err := db.Debug().Model(&Task{}).Create(&p).Error

	if err != nil {
		return &Task{}, err
	}

	return p, nil
}