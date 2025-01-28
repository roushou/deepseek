package deepseek

import (
	"errors"

	"github.com/roushou/deepseek/internal/http_client"
)

const (
	DefaultBaseURL = "https://api.deepseek.com"
)

type Option func(opts *options) error

type options struct {
	baseURL string
}

func WithBaseURL(baseURL string) Option {
	return func(opts *options) error {
		if baseURL == "" {
			return errors.New("invalid base URL")
		}
		opts.baseURL = baseURL
		return nil
	}
}

type Client struct {
	BaseURL string
	Balance *BalancesClient
	Chats   *ChatsClient
	Models  *ModelsClient
}

func NewClient(apiKey string, opts ...Option) (*Client, error) {
	options := &options{baseURL: DefaultBaseURL}
	for _, opt := range opts {
		if err := opt(options); err != nil {
			return nil, err
		}
	}

	httpClient, err := http_client.NewClient(options.baseURL)
	if err != nil {
		return nil, err
	}
	httpClient.SetHeader("Accept", "application/json")
	httpClient.SetHeader("Content-type", "application/json")
	httpClient.SetBearer(apiKey)

	return &Client{
		BaseURL: options.baseURL,
		Balance: &BalancesClient{httpClient},
		Chats:   &ChatsClient{httpClient},
		Models:  &ModelsClient{httpClient},
	}, nil
}
