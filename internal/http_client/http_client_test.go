package http_client_test

import (
	"testing"

	"github.com/roushou/deepseek/internal/http_client"
)

func TestNewClient(t *testing.T) {
	client, _ := http_client.NewClient("https://example.com")
	if client.BaseURL != "https://example.com" {
		t.Errorf("NewClient() BaseURL not set correctly. Got %s, want %s", client.BaseURL, "https://example.com")
	}
}

func TestSetBaseURL(t *testing.T) {
	client, _ := http_client.NewClient("http://old.example.com")
	client.SetBaseURL("http://new.example.com")
	if client.BaseURL != "http://new.example.com" {
		t.Errorf("SetBaseURL did not change BaseURL, got %s", client.BaseURL)
	}
}

func TestSetHeader(t *testing.T) {
	client, _ := http_client.NewClient("")
	client.SetHeader("Content-Type", "application/json")
	if client.Header.Get("Content-Type") != "application/json" {
		t.Errorf("Header not set correctly, got %s", client.Header.Get("Content-Type"))
	}
}

func TestSetHeaders(t *testing.T) {
	client, _ := http_client.NewClient("")
	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
	client.SetHeaders(headers)
	if client.Header.Get("Content-Type") != "application/json" || client.Header.Get("Accept") != "application/json" {
		t.Errorf("Headers not set correctly")
	}
}

func TestSetBearer(t *testing.T) {
	client, _ := http_client.NewClient("")
	client.SetBearer("token123")
	expected := "Bearer token123"
	if client.Header.Get("Authorization") != expected {
		t.Errorf("Authorization header not set correctly, got %s", client.Header.Get("Authorization"))
	}
}

func TestNewRequest(t *testing.T) {
	client, _ := http_client.NewClient("http://example.com")
	req, err := client.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Errorf("NewRequest returned an error: %v", err)
	}
	if req.URL.String() != "http://example.com/test" {
		t.Errorf("Request URL incorrect, got %s", req.URL.String())
	}
	if req.Method != "GET" {
		t.Errorf("Request method incorrect, got %s", req.Method)
	}
}
