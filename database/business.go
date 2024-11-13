package database

import (
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
)

type Business struct {
	GUID           uuid.UUID `gorm:"primary_key" json:"guid"`
	BusinessUid    string    `json:"business_uid"`
	DepositNotify  string    `json:"deposit_notify"`
	WithdrawNotify string    `json:"withdraw_notify"`
	TxFlowNotify   string    `json:"tx_flow_notify"`
	Timestamp      uint64
}

type BusinessView interface {
	QueryBusinessByUid(string) (*Business, error)
}

type BusinessDb interface {
	BusinessView
	StorageBusiness(*Business) error
}

type businessDb struct {
	gorm *gorm.DB
}

func NewBusinessDb(db *gorm.DB) BusinessDb {
	return &businessDb{gorm: db}
}

func (db *businessDb) QueryBusinessByUid(businessUid string) (*Business, error) {
	var business *Business
	result := db.gorm.Table("business").Where("business_uid", businessUid).First(&business)
	if result.Error != nil {
		log.Error("query business all fail", "Err", result.Error)
		return nil, result.Error
	}
	return business, nil
}

func (db *businessDb) StorageBusiness(business *Business) error {
	result := db.gorm.Table("business").Create(business)
	return result.Error
}
