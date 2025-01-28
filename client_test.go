package deepseek_test

import (
	"testing"

	"github.com/roushou/deepseek"
)

func TestNewClient(t *testing.T) {
	testCases := []struct {
		name            string
		apiKey          string
		opts            []deepseek.Option
		expectedBaseURL string
		wantErr         bool
	}{
		{
			name:            "Default Base URL",
			apiKey:          "api-key",
			opts:            []deepseek.Option{},
			expectedBaseURL: deepseek.DefaultBaseURL,
			wantErr:         false,
		},
		{
			name:            "Custom base URL",
			apiKey:          "api-key",
			opts:            []deepseek.Option{deepseek.WithBaseURL("https://deepsoup.com")},
			expectedBaseURL: "https://deepsoup.com",
			wantErr:         false,
		},
		{
			name:            "Empty base URL in WithBaseURL() Option",
			apiKey:          "api-key",
			opts:            []deepseek.Option{deepseek.WithBaseURL("")},
			expectedBaseURL: "",
			wantErr:         true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client, err := deepseek.NewClient(testCase.apiKey, testCase.opts...)
			hasErr := err != nil
			if hasErr != testCase.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !testCase.wantErr && client != nil {
				if client.Balance == nil || client.Chats == nil || client.Models == nil {
					t.Errorf("NewClient() client components not properly initialized")
				}
				if client.BaseURL != testCase.expectedBaseURL {
					t.Errorf("NewClient() BaseURL = %v, want %v", client.BaseURL, testCase.expectedBaseURL)
				}
			}
		})
	}
}

func TestWithBaseURL(t *testing.T) {
	testCases := []struct {
		name     string
		baseURL  string
		wantErr  bool
		expected string
	}{
		{
			name:     "Valid URL",
			baseURL:  "https://example.com",
			wantErr:  false,
			expected: "https://example.com",
		},
		{
			name:     "Empty URL",
			baseURL:  "",
			wantErr:  true,
			expected: deepseek.DefaultBaseURL, // error case, no change in baseURL
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			client, err := deepseek.NewClient("api-key", deepseek.WithBaseURL(testCase.baseURL))
			hasErr := err != nil
			if hasErr != testCase.wantErr {
				t.Errorf("WithBaseURL() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if !hasErr && client.BaseURL != testCase.expected {
				t.Errorf("WithBaseURL() got = %v, want %v", client.BaseURL, testCase.baseURL)
				return
			}
		})
	}
}
