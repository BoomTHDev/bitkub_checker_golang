package service

import (
	"github.com/boomthdev/wld_check_bk/pkg/custom"
	_walletModel "github.com/boomthdev/wld_check_bk/pkg/wallet/model"
)

type WalletService interface {
	GetWallet(coin, symbol, apiKey, apiSecret string) (*_walletModel.WalletResponse, *custom.AppError)
}
