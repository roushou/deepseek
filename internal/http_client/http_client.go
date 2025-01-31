package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL string
	Header  http.Header

	httpClient *http.Client
}

// NewClient creates a new HTTP client with default settings and optional configurations.
func NewClient(baseURL string) (*Client, error) {
	return &Client{
		BaseURL:    baseURL,
		httpClient: http.DefaultClient,
		Header:     http.Header{},
	}, nil
}

// SetBaseURL method sets the base URL for the client instance.
func (c *Client) SetBaseURL(baseURL string) {
	c.BaseURL = baseURL
}

// SetHeader method sets a single header field and its value in the client instance.
// These headers will be applied to all requests from this client instance.
func (c *Client) SetHeader(key, value string) {
	c.Header.Set(key, value)
}

// SetHeaders method sets multiple header fields and their values at one go in the client instance.
// These headers will be applied to all requests from this client instance.
func (c *Client) SetHeaders(headers map[string]string) {
	for k, v := range headers {
		c.Header.Set(k, v)
	}
}

// SetBearer method sets the bearer token to the authorization header. This header will be applied to all requests from this client instance.
func (c *Client) SetBearer(token string) {
	c.Header.Set("Authorization", "Bearer "+token)
}

// NewRequest method constructs a new HTTP request.
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	url := c.BaseURL + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header = c.Header
	return req, err
}

func (c *Client) Do(req *http.Request, out interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if out == nil {
		return resp, nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK: // 200
		if out != nil {
			err := json.NewDecoder(bytes.NewReader(body)).Decode(out)
			if err != nil {
				return nil, err
			}
			return resp, nil
		}
	case http.StatusBadRequest: // 400
		return nil, fmt.Errorf("%w: %s", ErrInvalidFormat, string(body))
	case http.StatusUnauthorized: // 401
		return nil, fmt.Errorf("%w: %s", ErrAuthenticationFailed, string(body))
	case http.StatusPaymentRequired: // 402
		return nil, fmt.Errorf("%w: %s", ErrInsufficientBalance, string(body))
	case http.StatusUnprocessableEntity: // 422
		return nil, fmt.Errorf("%w: %s", ErrInvalidParameters, string(body))
	case http.StatusTooManyRequests: // 429
		return nil, fmt.Errorf("%w: %s", ErrRateLimitExceeded, string(body))
	case http.StatusInternalServerError: // 500
		return nil, fmt.Errorf("%w: %s", ErrServer, string(body))
	case http.StatusServiceUnavailable: // 503
		return nil, fmt.Errorf("%w: %s", ErrServiceUnavailable, string(body))
	default:
		return nil, fmt.Errorf("unexpected HTTP status %d: %s", resp.StatusCode, string(body))
	}

	return nil, fmt.Errorf("unexpected HTTP status %d", resp.StatusCode)
}
