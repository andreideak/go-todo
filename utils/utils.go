package utils

import "os"

func GetNetwork() string {
	return os.Getenv("bsc_testnet")
}

func GetPrivateKey(AccountType string) string {
	prvKey := ""
	switch AccountType == "prod" {
	case true:
		prvKey = os.Getenv("privateKeyProd")
	case false:
		prvKey = os.Getenv("privateKeyDev")
	}
	return prvKey
}

func SetPrivateKey(accType, prvKey string) {
	switch accType == "prod" {
	case true:
		os.Setenv("privateKeyProd", prvKey)
	case false:
		os.Setenv("privateKeyDev", prvKey)

	}
}
