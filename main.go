package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

// WalletInfo represents the structure of wallet information to be saved
type WalletInfo struct {
	Seed        string `json:"seed"`
	AddressV4R2 string `json:"addressV4R2"`
}

func main() {


	client := liteclient.NewConnectionPool()

	configUrl := "https://ton-blockchain.github.io/global.config.json"


	err := client.AddConnectionsFromConfigUrl(context.Background(), configUrl)
	if err != nil {
		panic(err)
	}
	api := ton.NewAPIClient(client)

	var numWallets int
	fmt.Print("Enter the number of wallets to create: ")
	_, err = fmt.Scan(&numWallets)
	if err != nil {
		log.Fatal("Failed to read input:", err)
		return
	}

	done := make(chan struct{})

	var walletInfos []WalletInfo

	for i := 0; i < numWallets; i++ {
		go func() {
			defer func() {
				done <- struct{}{}
			}()

			seed := wallet.NewSeed()
			seedString := strings.Join(seed, " ") // Joining the slice into a single string
			fmt.Println("Seed phrase:", seedString)

			w, err := wallet.FromSeed(api, seed, wallet.V4R2) // Passing seed as slice of strings
			if err != nil {
				log.Println("FromSeed err:", err)
				return
			}

			address := w.Address().String()

			// Create WalletInfo struct
			walletInfo := WalletInfo{
				Seed:        seedString,
				AddressV4R2: address,
			}

			// Append WalletInfo to slice
			walletInfos = append(walletInfos, walletInfo)

			fmt.Println("Wallet address:", address)

			//block, err := api.CurrentMasterchainInfo(context.Background())
			if err != nil {
				log.Println("CurrentMasterchainInfo err:", err)
				return
			}

			// balance, err := w.GetBalance(context.Background(), block)
			// if err != nil {
			// 	log.Println("GetBalance err:", err)
			// 	return
			// }

			// fmt.Println("Balance:", balance)
		}()
	}

	for i := 0; i < numWallets; i++ {
		<-done
	}

	// Save wallet information to JSON file
	file, err := os.Create("wallets.json")
	if err != nil {
		log.Fatal("Failed to create file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(walletInfos)
	if err != nil {
		log.Fatal("Failed to encode wallets to JSON:", err)
		return
	}

	fmt.Println("Wallets information saved to wallets.json")
}
