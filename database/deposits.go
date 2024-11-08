package database

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

type Deposits struct {
	GUID         uuid.UUID      `gorm:"primaryKey" json:"guid"`
	BlockHash    common.Hash    `gorm:"column:block_hash;serializer:bytes"  db:"block_hash" json:"block_hash"`
	BlockNumber  *big.Int       `gorm:"serializer:u256;column:block_number" db:"block_number" json:"BlockNumber" form:"block_number"`
	Hash         common.Hash    `gorm:"column:hash;serializer:bytes"  db:"hash" json:"hash"`
	FromAddress  common.Address `json:"from_address" gorm:"serializer:bytes;column:from_address"`
	ToAddress    common.Address `json:"to_address" gorm:"serializer:bytes;column:to_address"`
	TokenAddress common.Address `json:"token_address" gorm:"serializer:bytes;column:token_address"`
	TokenId      string         `json:"token_id" gorm:"column:token_id"`
	TokenMeta    string         `json:"token_meta" gorm:"column:token_meta"`
	Fee          *big.Int       `gorm:"serializer:u256;column:fee" db:"fee" json:"Fee" form:"fee"`
	Amount       *big.Int       `gorm:"serializer:u256;column:amount" db:"amount" json:"Amount" form:"amount"`
	Status       uint8          `json:"status"` // 0:充值确认中,1:充值钱包层已到账；2:充值已通知业务层；3:充值完成
	Timestamp    uint64
}

type DepositsView interface {
}

type DepositsDB interface {
	DepositsView

	StoreDeposits(string, []Deposits, uint64) error
	UpdateDepositsStatus(requestId string, blockNumber uint64) error
}

type depositsDB struct {
	gorm *gorm.DB
}

func (db *depositsDB) UpdateDepositsStatus(requestId string, blockNumber uint64) error {
	result := db.gorm.Table("deposits_"+requestId).Where("status = ? and block_number <= ?", 0, blockNumber).Updates(map[string]interface{}{"status": 1})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil
		}
		return result.Error
	}
	return nil
}

func NewDepositsDB(db *gorm.DB) DepositsDB {
	return &depositsDB{gorm: db}
}

func (db *depositsDB) StoreDeposits(requestId string, depositList []Deposits, depositLength uint64) error {
	result := db.gorm.Table("deposits_"+requestId).CreateInBatches(&depositList, int(depositLength))
	if result.Error != nil {
		log.Error("create deposit batch fail", "Err", result.Error)
		return result.Error
	}
	return nil
}
