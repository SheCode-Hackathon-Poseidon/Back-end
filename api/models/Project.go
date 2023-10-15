package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Post is...
type Project struct {
    gorm.Model
	UserId 		int	`json:"user_id"`
    Status      string `gorm:"type:project_status" json:"status"`
    Name        string `gorm:"type:varchar" json:"name"`
    Description string `gorm:"type:varchar" json:"description"`
    ShareMode   string `gorm:"type:share_mode" json:"share_mode"`
	Tasks       []Task `json:"tasks"`
}

// Prepare is...
func (p *Project) Prepare() {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate is...
func (p *Project) Validate() error {
	
	return nil
}

// Create is...
func (p *Project) Create(db *gorm.DB) (*Project, error) {
	err := db.Debug().Model(&Project{}).Create(&p).Error

	if err != nil {
		return &Project{}, err
	}

	return p, nil
}

// FindAllPosts is...
func (p *Project) FindAllPosts(db *gorm.DB) (*[]Project, error) {
	var err error
	posts := []Project{}

	err = db.Debug().Model(&Project{}).Limit(100).Find(&posts).Error

	if err != nil {
		return &[]Project{}, err
	}

	// if len(posts) > 0 {
	// 	for i := range posts {
	// 		_err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error

	// 		if _err != nil {
	// 			return &[]Project{}, _err
	// 		}
	// 	}
	// }

	return &posts, nil
}

// FindPostByID is...
func (p *Project) FindPostByID(db *gorm.DB, pid uint32) (*Project, error) {
	var err error

	err = db.Debug().Model(&Project{}).Where("id = ?", pid).Take(&p).Error

	if err != nil {
		return &Project{}, err
	}

	return p, nil
}

// UpdatePost is...
func (p *Project) UpdatePost(db *gorm.DB) (*Project, error) {
	// var err error

	// err = db.Debug().Model(&Project{}).Where("id = ?", p.ID).Updates(Post{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error

	// if err != nil {
	// 	return &Project{}, err
	// }

	// if p.ID != 0 {
	// 	err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error

	// 	if err != nil {
	// 		return &Project{}, err
	// 	}
	// }

	return p, nil
}

// DeletePost is...
func (p *Project) DeletePost(db *gorm.DB, pid uint32, uid uint32) (int64, error) {
	db = db.Debug().Model(&Project{}).Where("id = ? and author_id = ?", pid, uid).Take(&Project{}).Delete(&Project{})

	// if db.Error != nil {
	// 	if gorm.IsRecordNotFoundError(db.Error) {
	// 		return 0, errors.New("Post not found")
	// 	}

	// 	return 0, db.Error
	// }

	return db.RowsAffected, nil
}
