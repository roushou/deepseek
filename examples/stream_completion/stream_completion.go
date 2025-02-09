package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/roushou/deepseek"
)

func main() {
	client, err := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	stream := client.Chats.CreateStreamCompletion(ctx, deepseek.StreamCompletionArgs{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.Message{
			{
				Role:    deepseek.SystemRole,
				Content: "You are a helpful assistant",
			},
			{
				Role:    deepseek.UserRole,
				Content: "Explain Fermat's last theorem",
			},
		},
	})
	defer stream.Close()

	for stream.Next() {
		fmt.Print(stream.Current().Choices[0].Delta.Content)
	}
}
