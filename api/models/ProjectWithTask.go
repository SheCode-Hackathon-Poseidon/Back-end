package models

// import (
// 	"time"

// 	"github.com/jinzhu/gorm"
// )

// // Post is...
// type ProjectWithTask struct {
//     gorm.Model
// 	UserId 		int
//     Status      string `gorm:"type:project_status" json:"status"`
//     Name        string `gorm:"type:varchar" json:"name"`
//     Description string `gorm:"type:varchar" json:"description"`
//     ShareMode   string `gorm:"type:share_mode" json:"share_mode"`
// 	Tasks		[]Task 
// }

// // Prepare is...
// func (p *ProjectWithTask) Prepare() {
// 	p.CreatedAt = time.Now()
// 	p.UpdatedAt = time.Now()
// }

// // Create is...
// func (p *ProjectWithTask) Create(db *gorm.DB) (*ProjectWithTask, error) {
// 	var err error

// 	err = db.Debug().Model(&Project{}).Create(&p).Error

// 	if err != nil {
// 		return &Project{}, err
// 	}

// 	return p, nil
// }