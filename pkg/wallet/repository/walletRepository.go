package repository

type WalletRepository interface {
	GetServerTime() ([]byte, error)
	GetCoinPrice(timestamp, apiKey, apiSecret string, b []byte) ([]byte, error)
	GetLastPrice(symbol string) ([]byte, error)
}
