package main

import (
	"fmt"
	"log"
	"os"

	"github.com/roushou/deepseek"
)

func main() {
	client, err := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	balance, err := client.Balance.GetUserBalance()
	if err != nil {
		log.Fatalf("failed to get user balance: %v", err)
	}

	fmt.Printf("Available: %t\n", balance.IsAvailable)
	for _, info := range balance.BalanceInfos {
		fmt.Printf("Currency: %s\n", info.Currency)
		fmt.Printf("Total by: %s\n", info.TotalBalance)
		fmt.Printf("Granted by: %s\n", info.GrantedBalance)
		fmt.Printf("Topped up: %s\n", info.ToppedUpBalance)
		fmt.Println("==================")
	}
}
