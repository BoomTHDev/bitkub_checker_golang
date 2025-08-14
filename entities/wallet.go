package entities

type Wallet struct {
	Error  int            `json:"error"`
	Result map[string]any `json:"result"`
}
