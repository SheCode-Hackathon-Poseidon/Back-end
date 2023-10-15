package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email       string `gorm:"type:varchar;unique_index"`
	Password    string `gorm:"type:varchar"`
	FullName    string `gorm:"type:varchar"`
	PhoneNumber string `gorm:"type:char(20)"`
	City        string `gorm:"type:varchar"`
	Country     string `gorm:"type:varchar"`
	Token       string `gorm:"type:varchar"`
	Dob         time.Time
	TotalPoint  int
	AvatarURL   string `gorm:"type:varchar"`
}

// Hash is...
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword is...
func VerifyPassword(sent_password, password string) bool {
	return sent_password == password
}

// BeforeSave is...
func (u *User) BeforeSave() error {
	hashedPassowrd, err := Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassowrd)

	return nil
}

// Prepare is...
func (u *User) Prepare() error {
	u.ID = 0
	u.FullName = html.EscapeString(strings.TrimSpace(u.FullName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return nil
}

// Validate is...
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.FullName == "" {
			return errors.New("FullName required")
		}

		if u.Password == "" {
			return errors.New("Password required")
		}

		if u.Email == "" {
			return errors.New("Email required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email format invalid")
		}

		return nil

	case "login":
		if u.Password == "" {
			return errors.New("Password required")
		}

		if u.Email == "" {
			return errors.New("Email required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email format invalid")
		}

		return nil

	default:
		if u.FullName == "" {
			return errors.New("FullName required")
		}

		if u.Password == "" {
			return errors.New("Password required")
		}

		if u.Email == "" {
			return errors.New("Email required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email format invalid")
		}

		return nil
	}
}

// SaveUser is...
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// FindAllUsers is...
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}

	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

// FindUserByID is...
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// FindUserByToken is...
func (u *User) FindUserByToken(db *gorm.DB, token string) (*User, error) {
	err := db.Debug().Model(User{}).Where("token = ?", token).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

// UpdateUser is...
func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave()

	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).UpdateColumns(
		map[string]interface{}{
			"password":   u.Password,
			"full_name":  u.FullName,
			"email":      u.Email,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return &User{}, db.Error
	}

	return u, nil
}

// DeleteUser is...
func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
