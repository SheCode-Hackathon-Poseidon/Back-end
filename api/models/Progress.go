package models

import (
	"github.com/jinzhu/gorm"
)

type Progress struct {
	gorm.Model
	ProjectID   int	`gorm:"type:varchar" json:"project_id"`
	UserID      int `gorm:"type:varchar" json:"user_id"`
	Description string `gorm:"type:varchar" json:"description"`
}

func (p *Progress) Create(db *gorm.DB) (*Progress, error) {
	err := db.Debug().Model(&Progress{}).Create(&p).Error

	if err != nil {
		return &Progress{}, err
	}

	return p, nil
}

func (p *Progress) FindProgressOfProject(db *gorm.DB, pid uint32) ([]Progress, error) {
    var progress []Progress
    err := db.Debug().Model(&Progress{}).Where("project_id = ?", pid).Find(&progress).Error

    if err != nil {
        return []Progress{}, err
    }

    return progress, nil
}