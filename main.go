package main

import (
	"log"
	"os"

	"github.com/andreideak/go-todo.git/binance"
	"github.com/joho/godotenv"
)

// Environment variables
// TODO: replace with ENV variables
var testnet = "testnet-dex.binance.org"

// var mainnet = "dex.binance.org"

func init() {
	// Load environment variables from .ENV file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load environment variables .env file due to exception:\n%v\n", err)
	} else {
		log.Println("Successfully loaded environment variables")
	}

}

func main() {
	// Retrieve raw private key from .env file
	privateKey := os.Getenv("PRV_KEY_TESTNET")
	// privateKey := os.Getenv("PRV_KEY_MAINNET")

	// Create 'quit' channel (used for Subscribption to Binance)
	quit := make(chan struct{})
	defer close(quit)
	// Create newBlock channel (for retrieving current block height)
	chanNewBlock := make(chan int64)
	defer close(chanNewBlock)

	// Initialize Key Manager
	keyManager := binance.InitializeKeyManager(privateKey)

	// Initialize SDK
	client := binance.InitializeSDK(keyManager, testnet)

	// Subscribe to new Block Height
	go binance.GetNewBlockHeight(client, quit, chanNewBlock)

	for item := range chanNewBlock {
		log.Printf("New block - #%v", item)
	}

}
