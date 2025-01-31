# DeepSeek Go

`deepseek` is a Go SDK for the [DeepSeek API](https://api-docs.deepseek.com/).

## Installation

```bash
go get github.com/roushou/deepseek
```

## Usage

See [examples](./examples) for more.

```go
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

	completion, err := client.Chats.CreateCompletion(deepseek.CompletionArgs{
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
```

## License

This project is licensed under the MIT License. See the [License](./LICENSE) file for details.
