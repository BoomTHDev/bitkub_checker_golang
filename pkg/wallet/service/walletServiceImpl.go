package service

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/boomthdev/wld_check_bk/entities"
	"github.com/boomthdev/wld_check_bk/pkg/custom"
	_walletModel "github.com/boomthdev/wld_check_bk/pkg/wallet/model"
	_walletRepository "github.com/boomthdev/wld_check_bk/pkg/wallet/repository"
)

type walletServiceImpl struct {
	walletRepository _walletRepository.WalletRepository
}

func NewWalletService(walletRepository _walletRepository.WalletRepository) WalletService {
	return &walletServiceImpl{walletRepository: walletRepository}
}

func (s *walletServiceImpl) GetWallet(coin, symbol, apiKey, apiSecret string) (*_walletModel.WalletResponse, *custom.AppError) {
	var (
		b      []byte
		err    error
		ts     int64
		wallet entities.Wallet
		all    map[string]map[string]any
		msg    string
	)

	b, err = s.walletRepository.GetServerTime()
	if err != nil {
		return nil, custom.ErrIntervalServer("Failed to get server time", err)
	}

	if err = json.Unmarshal(b, &ts); err != nil {
		return nil, custom.ErrIntervalServer("Failed to unmarchal json", err)
	}
	timestamp := strconv.FormatInt(ts, 10)
	b, err = s.walletRepository.GetCoinPrice(timestamp, apiKey, apiSecret, b)
	if err != nil {
		return nil, custom.ErrIntervalServer("Failed to get coin price", err)
	}

	if err = json.Unmarshal(b, &wallet); err != nil {
		return nil, custom.ErrIntervalServer("Failed to unmarshal json", err)
	}

	val, ok := wallet.Result[coin].(float64)
	if !ok {
		msg = fmt.Sprintf("%s not found", coin)
		return nil, custom.ErrNotFound(msg, nil)
	}

	b, err = s.walletRepository.GetLastPrice(symbol)
	if err != nil {
		return nil, custom.ErrIntervalServer("Failed to get last price", err)
	}

	if err = json.Unmarshal(b, &all); err != nil {
		return nil, custom.ErrIntervalServer("Failed to unmarshal json", err)
	}

	row, ok := all[symbol]
	if !ok {
		msg = fmt.Sprintf("Symbol %s not found in ticker", symbol)
		return nil, custom.ErrNotFound(msg, nil)
	}

	last, ok := row["last"].(float64)
	if !ok {
		msg = fmt.Sprintf("Symbol %s last price not float", symbol)
		return nil, custom.ErrNotFound(msg, nil)
	}

	valueTHB := val * last

	resp := &_walletModel.WalletResponse{
		CoinExists:   val,
		THBCoinPrice: last,
		MyTHB:        valueTHB,
	}

	return resp, nil
}
