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

	models, err := client.Models.ListModels()
	if err != nil {
		log.Fatalf("failed to get models: %v", err)
	}

	for _, model := range models.Data {
		fmt.Printf("ID: %s\n", model.ID)
		fmt.Printf("Object: %s\n", model.Object)
		fmt.Printf("Owned by: %s\n", model.OwnedBy)
		fmt.Println("==================")
	}

	model, err := client.Models.GetModel("deepseek-chat")
	if err != nil {
		log.Fatalf("failed to get model 'deepseek-chat'")
	}
	fmt.Printf("ID: %s\n", model.ID)
	fmt.Printf("Object: %s\n", model.Object)
	fmt.Printf("Owned by: %s\n", model.OwnedBy)
	fmt.Println("==================")
}
