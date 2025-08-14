package repository

type WalletRepository interface {
	GetServerTime() ([]byte, error)
	GetCoinPrice(timestamp string, b []byte) ([]byte, error)
	GetLastPrice(symbol string) ([]byte, error)
}
