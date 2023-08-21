package mutualfund

import (
	"time"

	"github.com/samar2170/portfolio-manager-v4/pkg/db"
	"gorm.io/gorm"
)

type MutualFund struct {
	*gorm.Model
	ID                   int
	SchemeName           string `gorm:"index"`
	SchemeCategory       string `gorm:"index"`
	SchemeNavName        string `gorm:"index"`
	ParentSchemeCategory string `gorm:"index"`
	PriceToBeUpdated     bool
}

type MutualFundNavHistory struct {
	*gorm.Model
	ID           int
	MutualFund   MutualFund `gorm:"foreignKey:MutualFundID;index"`
	MutualFundID int
	Nav          float64
	Date         time.Time
	Source       string
}

func (m *MutualFundNavHistory) Create() error {
	err := db.DB.Create(&m).Error
	return err
}

func (m *MutualFund) Create() error {
	err := db.DB.Create(&m).Error
	return err
}

func (m *MutualFund) GetOrCreate() (MutualFund, error) {
	err := db.DB.FirstOrCreate(&m, MutualFund{SchemeName: m.SchemeName, SchemeNavName: m.SchemeNavName}).Error
	return *m, err
}
