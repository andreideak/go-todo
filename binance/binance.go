package binance

import (
	"log"

	sdk "github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/client/websocket"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
)

var testnet = "testnet-dex.binance.org"
// var mainnet = "dex.binance.org"

// Struct to define the Binance response with the latest block height
type BlockHeightEvent struct {
	BlockHeight int64 `json:"h"`
}

// Function to initialize a Key Manager, from a raw Private Key
func InitializeKeyManager(prvKey string) keys.KeyManager {
	keyManager, err := keys.NewPrivateKeyManager(prvKey)
	if err != nil {
		log.Fatalf("Unable to generate keyManager from Raw Private Key")
	} else {
		log.Println("Successfully generated keyManager from Raw Private Key")
	}

	return keyManager
}

// Function to initialize the SDK, based on the previously initialized Key Manager
func InitializeSDK(keyManager keys.KeyManager, network string) sdk.DexClient {
	client, err := sdk.NewDexClient(network, types.TestNetwork, keyManager)
	if err != nil {
		log.Fatalf("Unable to connect to Binance Smart Chain due to exception:\n%v\n", err)
	} else {
		log.Printf("Successfully connected to %v\n", network)
	}

	return client
}

// Function to retrieve the latest block height
func GetNewBlockHeight(client sdk.DexClient, quit chan struct{}, chanNewBlock chan int64) {
	// Subscribe for new block
	err := client.SubscribeBlockHeightEvent(quit, func(event *websocket.BlockHeightEvent) {
		chanNewBlock <- event.BlockHeight
	}, func(err error) {
	}, nil)
	if err != nil {
		log.Fatalf("Unable to subscribe for new block height events due to exception:\n%v\n", err)
	} else {
		log.Println("Successfully subscribed for new block height events")
	}
}
