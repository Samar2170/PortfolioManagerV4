package portfolio

import (
	"github.com/samar2170/portfolio-manager-v4/internal/models"
	"github.com/samar2170/portfolio-manager-v4/pkg/db"
	"github.com/samar2170/portfolio-manager-v4/security/bond"
	mutualfund "github.com/samar2170/portfolio-manager-v4/security/mutual-fund"
	"github.com/samar2170/portfolio-manager-v4/security/stock"
	"gorm.io/gorm"
)

type HoldingInterface interface {
	create() error
	getInvestedValue() float64
	// getCurrentValue() float64
	// getProfit() float64
	// getProfitPercentage() float64
}

type StockHolding struct {
	*gorm.Model
	StockID  int
	Stock    *stock.Stock
	Quantity int
	BuyPrice float64
	Account  models.DematAccount
}

type BondHolding struct {
	*gorm.Model
	BondID   int
	Bond     *bond.Bond
	Quantity int
	BuyPrice float64
	Account  models.DematAccount
}

type MutualFundHolding struct {
	*gorm.Model
	MutualFundID int
	MutualFund   *mutualfund.MutualFund
	Quantity     int
	BuyPrice     float64
	Account      models.DematAccount
}

type ETSHolding struct {
	*gorm.Model
	ETSID    int
	ETS      *stock.Stock
	Quantity int
	BuyPrice float64
	Account  models.DematAccount
}

func (s *StockHolding) create() error {
	return db.DB.Create(s).Error
}

func (b *BondHolding) create() error {
	return db.DB.Create(b).Error
}
func (mf *MutualFundHolding) create() error {
	return db.DB.Create(mf).Error
}

func (e *ETSHolding) create() error {
	return db.DB.Create(e).Error
}

func (s *StockHolding) getInvestedValue() float64 {
	return float64(s.Quantity) * s.BuyPrice
}

func (b *BondHolding) getInvestedValue() float64 {
	return float64(b.Quantity) * b.BuyPrice
}

func (mf *MutualFundHolding) getInvestedValue() float64 {
	return float64(mf.Quantity) * mf.BuyPrice
}

func (e *ETSHolding) getInvestedValue() float64 {
	return float64(e.Quantity) * e.BuyPrice
}
