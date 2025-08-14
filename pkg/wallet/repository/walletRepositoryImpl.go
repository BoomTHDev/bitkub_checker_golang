package repository

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/boomthdev/wld_check_bk/config"
	"github.com/boomthdev/wld_check_bk/pkg/util"
)

type walletRepositoryImpl struct {
	bitkub *config.Bitkub
}

func NewWalletRepositoryImpl(bitkub *config.Bitkub) WalletRepository {
	return &walletRepositoryImpl{bitkub: bitkub}
}

const baseURL = "https://api.bitkub.com"

func (r *walletRepositoryImpl) GetServerTime() ([]byte, error) {
	endpoint := "/api/v3/servertime"
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return body, nil
}

func (r *walletRepositoryImpl) GetCoinPrice(timestamp string, b []byte) ([]byte, error) {
	endpoint := "/api/v3/market/wallet"

	sig := util.GenerateSignature(r.bitkub.BitkubApiSecret, timestamp, http.MethodPost, endpoint, string(b))

	req, err := http.NewRequest(http.MethodPost, baseURL+endpoint, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-BTK-APIKEY", r.bitkub.BitkubApiKey)
	req.Header.Set("X-BTK-TIMESTAMP", timestamp)
	req.Header.Set("X-BTK-SIGN", sig)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)

	return body, nil
}

func (r *walletRepositoryImpl) GetLastPrice(symbol string) ([]byte, error) {
	endpoint := "/api/market/ticker"
	resp, err := http.Get(baseURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	return body, nil
}
