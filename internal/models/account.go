package models

import (
	"github.com/samar2170/portfolio-manager-v4/pkg/db"
	"gorm.io/gorm"
)

type BankAccount struct {
	*gorm.Model
	ID        int
	UserCID   string `gorm:"index"`
	Bank      string
	AccountNo string `gorm:"unique;index"`
}

type DematAccount struct {
	*gorm.Model
	ID          int
	UserCID     string `gorm:"index"`
	AccountCode string `gorm:"unique;index"`
	Broker      string
}

type GeneralAccount struct {
	*gorm.Model
	ID          int
	UserCID     string `gorm:"unique;index"`
	AccountCode string `gorm:"unique;index"`
}

func (b *BankAccount) create() error {
	return db.DB.Create(b).Error
}
func (b *BankAccount) update() error {
	return db.DB.Save(b).Error
}

func (d *DematAccount) create() error {
	return db.DB.Create(d).Error
}

func (d *DematAccount) update() error {
	return db.DB.Save(d).Error
}

func (g *GeneralAccount) create() error {
	return db.DB.Create(g).Error
}

func (g *GeneralAccount) update() error {
	return db.DB.Save(g).Error
}
