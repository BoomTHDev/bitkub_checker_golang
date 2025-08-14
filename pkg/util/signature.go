package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSignature(secret, timestamp, method, endpoint, payload string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(timestamp + method + endpoint + payload))
	return hex.EncodeToString(mac.Sum(nil))
}
