package models

import (
	"github.com/samar2170/portfolio-manager-v4/pkg/db"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID       int
	Username string `gorm:"unique"`
	UserCID  string `gorm:"unique"`
}

// User signs up/login, we call rpc, save CID and username,
func (u *User) create() error {
	return db.DB.Create(u).Error
}

func (u *User) update() error {
	return db.DB.Save(u).Error
}

type Watchlist struct {
	*gorm.Model
	UserCID string `gorm:"index"`
	Name    string
}
