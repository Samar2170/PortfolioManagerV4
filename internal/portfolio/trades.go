package portfolio

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"reflect"

	"github.com/samar2170/portfolio-manager-v4/internal/models"
	"github.com/samar2170/portfolio-manager-v4/pkg/db"
	"github.com/samar2170/portfolio-manager-v4/pkg/utils/structs"
	"github.com/samar2170/portfolio-manager-v4/security/bond"
	"github.com/samar2170/portfolio-manager-v4/security/ets"
	mutualfund "github.com/samar2170/portfolio-manager-v4/security/mutual-fund"
	"github.com/samar2170/portfolio-manager-v4/security/stock"
	"gorm.io/gorm"
)

type BlockHash struct {
	*gorm.Model
	UserCID      string
	Hash         string
	PreviousHash string
}

func (b *BlockHash) create() error {
	return db.DB.Create(b).Error
}
func GetLatestBlockHash(userCID string) (BlockHash, error) {
	var b BlockHash
	err := db.DB.Where("user_cid = ?", userCID).Last(&b).Error
	return b, err
}

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
}

func (s *StockTrade) create() error {
	return db.DB.Create(s).Error
}
func (s *StockTrade) getAccount() models.DematAccount {
	return s.Account
}

func (b *BondTrade) create() error {

	return db.DB.Create(b).Error
}
func (b *BondTrade) getAccount() models.DematAccount {
	return b.Account
}

func (m *MutualFundTrade) create() error {
	return db.DB.Create(m).Error
}
func (m *MutualFundTrade) getAccount() models.DematAccount {
	return m.Account
}

func (e *ETSTrade) create() error {
	return db.DB.Create(e).Error
}
func (e *ETSTrade) getAccount() models.DematAccount {
	return e.Account
}

type TradeInterface interface {
	create() error
	getAccount() models.DematAccount
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
func RegisterTrade(td TradeInterface) {
	createModelInstance(td)
	createHashBlockForTrade(&td)
}

func createHashBlockForTrade(td *TradeInterface) error {
	s := structs.New(td)
	m := fmt.Sprint(s.Map())
	accountCID := (*td).getAccount().UserCID
	latestBlock, err := GetLatestBlockHash(accountCID)
	if err != nil {
		return err
	}
	hasher := sha1.New()
	hasher.Write([]byte(m))
	hash := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hash)
	blockHash := BlockHash{
		UserCID:      accountCID,
		Hash:         hashString,
		PreviousHash: latestBlock.Hash,
	}
	err = blockHash.create()
	if err != nil {
		return err
	}
	return nil
}
