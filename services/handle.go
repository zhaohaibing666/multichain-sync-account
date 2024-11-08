package services

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"

	"github.com/zhaohaibing666/multichain-sync-account/database"
	dal_wallet_go "github.com/zhaohaibing666/multichain-sync-account/protobuf/dal-wallet-go"
)

func (bws *BusinessMiddleWireServices) BusinessRegister(ctx context.Context, request *dal_wallet_go.BusinessRegisterRequest) (*dal_wallet_go.BusinessRegisterResponse, error) {
	if request.RequestId == "" || request.DepositNotify == "" || request.WithdrawNotify == "" {
		return &dal_wallet_go.BusinessRegisterResponse{
			Code: dal_wallet_go.ReturnCode_ERROR,
			Msg:  "invalid params",
		}, nil
	}
	business := &database.Business{
		GUID:           uuid.New(),
		BusinessUid:    request.RequestId,
		DepositNotify:  request.DepositNotify,
		WithdrawNotify: request.WithdrawNotify,
		TxFlowNotify:   request.TxFlowNotify,
		Timestamp:      uint64(time.Now().Unix()),
	}
	err := bws.db.Business.StoreBusiness(business)
	if err != nil {
		log.Error("store business fail", "err", err)
		return &dal_wallet_go.BusinessRegisterResponse{
			Code: dal_wallet_go.ReturnCode_ERROR,
			Msg:  "invalid params",
		}, nil
	}
	return &dal_wallet_go.BusinessRegisterResponse{
		Code: dal_wallet_go.ReturnCode_SUCCESS,
		Msg:  "config business success",
	}, nil
}
