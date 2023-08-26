package portfolio

import (
	"errors"
	"reflect"

	"github.com/samar2170/portfolio-manager-v4/internal/models"
	"github.com/samar2170/portfolio-manager-v4/pkg/db"
	"github.com/samar2170/portfolio-manager-v4/security/bond"
	"github.com/samar2170/portfolio-manager-v4/security/ets"
	mutualfund "github.com/samar2170/portfolio-manager-v4/security/mutual-fund"
	"github.com/samar2170/portfolio-manager-v4/security/stock"
	"gorm.io/gorm"
)

type StockTrade struct {
	*gorm.Model
	ID        int
	StockID   int
	Stock     *stock.Stock
	Quantity  int
	Price     float64
	TradeType string
	TradeDate string
	Account   models.DematAccount

	BlockHash         string
	PreviousBlockHash string
}

type BondTrade struct {
	*gorm.Model
	ID        int
	BondID    int
	Bond      *bond.Bond
	Quantity  int
	Price     float64
	TradeType string
	TradeDate string
	Account   models.DematAccount

	BlockHash string
}

type MutualFundTrade struct {
	*gorm.Model
	ID           int
	MutualFundID int
	MutualFund   *mutualfund.MutualFund
	Quantity     int
	Price        float64
	TradeType    string
	TradeDate    string
	Account      models.DematAccount

	BlockHash string
}

type ETSTrade struct {
	*gorm.Model
	ID        int
	ETSID     int
	ETS       *ets.ETS
	Quantity  int
	Price     float64
	TradeType string
	TradeDate string
	Account   models.DematAccount

	BlockHash string
}

func (s *StockTrade) create() error {
	return db.DB.Create(s).Error
}

func (b *BondTrade) create() error {
	return db.DB.Create(b).Error
}

func (m *MutualFundTrade) create() error {
	return db.DB.Create(m).Error
}

func (e *ETSTrade) create() error {
	return db.DB.Create(e).Error
}

type TradeInterface interface {
	create() error
}

func createModelInstance(td TradeInterface) error {
	err := td.create()
	modelName := reflect.TypeOf(td).Name()
	switch err {
	case nil:
		return nil
	case gorm.ErrDuplicatedKey:
		return errors.New("[gorm]:ErrDuplicatedKey" + modelName + " already exists")
	case gorm.ErrForeignKeyViolated:
		return errors.New("[gorm]:ErrForeignKeyViolated" + modelName + " already exists")
	default:
		return err
	}
}

// lets do it blockchain style
// func RegisterTrade(td TradeInterface) {

// }
