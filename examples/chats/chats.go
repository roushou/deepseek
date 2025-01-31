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

	completion, err := client.Chats.CreateCompletion(deepseek.ChatCompletionArgs{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.Message{
			{
				Role:    deepseek.SystemRole,
				Content: "You are a helpful assistant",
			},
			{
				Role:    deepseek.UserRole,
				Content: "Hello World",
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to create completion: %v", err)
	}

	fmt.Println(completion)
}
