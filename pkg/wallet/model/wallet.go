package model

type (
	WalletResponse struct {
		CoinExists   float64 `json:"coin_exists"`
		THBCoinPrice float64 `json:"thb_coin_price"`
		MyTHB        float64 `json:"my_thb"`
	}
)
