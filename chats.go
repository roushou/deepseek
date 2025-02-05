package deepseek

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/roushou/deepseek/internal/http_client"
	"github.com/roushou/deepseek/internal/ssestream"
)

type ChatsClient struct {
	httpClient *http_client.Client
}

func (c *ChatsClient) CreateCompletion(args CompletionArgs) (*CompletionResponse, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	req, err := c.httpClient.NewRequest(http.MethodPost, "/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var completion CompletionResponse
	_, err = c.httpClient.Do(req, &completion)
	return &completion, err
}

// CreateStreamCompletion streams chat completion using Server-Sent Events (SSE).
func (c *ChatsClient) CreateStreamCompletion(ctx context.Context, args StreamCompletionArgs) *ssestream.Stream[StreamCompletionChunk] {
	args.Stream = true

	body, err := json.Marshal(args)
	if err != nil {
		return nil
	}

	req, err := c.httpClient.NewRequest(http.MethodPost, "/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil
	}
	req.Header.Set("Accept", "text/event-stream")

	resp, err := c.httpClient.Do(req, nil)
	if err != nil {
		return nil
	}

	return ssestream.NewStream[StreamCompletionChunk](ssestream.NewDecoder(resp), nil)
}

// NewCompletionRequest creates a new chat completion request with default values.
func NewCompletionRequest(model ModelID) CompletionArgs {
	return CompletionArgs{
		Model:            model,
		Messages:         []Message{},
		FrequencyPenalty: 0,
		MaxTokens:        4096,
		PresencePenalty:  0,
		ResponseFormat:   &ResponseFormat{},
		Stop:             []string{},
		Temperature:      1,
		TopP:             1,
		Tools:            []Tool{},
		ToolChoice:       "",
		Logprobs:         false,
		TopLogprobs:      0,
	}
}

// NewCompletionRequest creates a new chat completion request with default values.
func NewStreamCompletionRequest(model ModelID) StreamCompletionArgs {
	return StreamCompletionArgs{
		Model:            model,
		Messages:         []Message{},
		FrequencyPenalty: 0,
		MaxTokens:        4096,
		PresencePenalty:  0,
		ResponseFormat:   &ResponseFormat{},
		Stream:           true,
		StreamOptions:    &StreamOptions{},
		Stop:             []string{},
		Temperature:      1,
		TopP:             1,
		Tools:            []Tool{},
		ToolChoice:       "",
		Logprobs:         false,
		TopLogprobs:      0,
	}
}

type CompletionArgs struct {
	// Model is the ID of the model to use.
	Model ModelID `json:"model"`

	// Messages contains the conversation history or context for the chat.
	Messages []Message `json:"messages"`

	// FrequencyPenalty adjusts the likelihood of repeating tokens based on their frequency in the text.
	//
	// Range: -2.0 to 2.0. Default: 0. Higher values decrease repetition.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// MaxTokens is the maximum number of tokens that can be generated in the chat completion. The total length of input tokens and generated tokens is limited by the model's context length.
	//
	// Integer between 1 and 8192. Defaults to 4096.
	MaxTokens int `json:"max_tokens,omitempty"`

	// PresencePenalty influences the model to introduce new topics by penalizing tokens based on their presence in the text.
	//
	// Range: -2.0 to 2.0. Default: 0. Higher values encourage new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// ResponseFormat defines the format of the response (e.g., text or JSON).
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`

	// Stop provides sequences at which to stop generating further tokens. Up to 16 sequences allowed.
	Stop []string `json:"stop,omitempty"`

	// Temperature controls the randomness of the output; higher values make it more creative, lower values more deterministic.
	//
	// Range: 0 to 2. Default: 1.
	Temperature float64 `json:"temperature,omitempty"`

	// TopP implements nucleus sampling where only tokens with cumulative probability up to this value are considered.
	//
	// For example, 0.1 means only the tokens comprising the top 10% probability mass are considered.
	TopP float64 `json:"top_p,omitempty"`

	// Tools lists functions that the model can call, limited to 128.
	Tools []Tool `json:"tools,omitempty"`

	// ToolChoice specifies which tool to use if tools are provided.
	ToolChoice string `json:"tool_choice,omitempty"`

	// Logprobs indicates if log probabilities should be returned for each token.
	Logprobs bool `json:"logprobs,omitempty"`

	// TopLogprobs specifies how many most likely tokens to return with their log probabilities; requires Logprobs to be true.
	//
	// Range: 0 to 20.
	TopLogprobs float64 `json:"top_logprobs,omitempty"`
}

type StreamCompletionArgs struct {
	// Model is the ID of the model to use.
	Model ModelID `json:"model"`

	// Messages contains the conversation history or context for the chat.
	Messages []Message `json:"messages"`

	// FrequencyPenalty adjusts the likelihood of repeating tokens based on their frequency in the text.
	//
	// Range: -2.0 to 2.0. Default: 0. Higher values decrease repetition.
	FrequencyPenalty float64 `json:"frequency_penalty,omitempty"`

	// MaxTokens is the maximum number of tokens that can be generated in the chat completion. The total length of input tokens and generated tokens is limited by the model's context length.
	//
	// Integer between 1 and 8192. Defaults to 4096.
	MaxTokens int `json:"max_tokens,omitempty"`

	// PresencePenalty influences the model to introduce new topics by penalizing tokens based on their presence in the text.
	//
	// Range: -2.0 to 2.0. Default: 0. Higher values encourage new topics.
	PresencePenalty float64 `json:"presence_penalty,omitempty"`

	// ResponseFormat defines the format of the response (e.g., text or JSON).
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`

	// Stop provides sequences at which to stop generating further tokens. Up to 16 sequences allowed.
	Stop []string `json:"stop,omitempty"`

	// Stream indicates whether to stream back partial results or return the full response at once.
	Stream bool `json:"stream,omitempty"`

	// StreamOptions configures whether to include usage statistics in the streaming response.
	StreamOptions *StreamOptions `json:"stream_options,omitempty"`

	// Temperature controls the randomness of the output; higher values make it more creative, lower values more deterministic.
	//
	// Range: 0 to 2. Default: 1.
	Temperature float64 `json:"temperature,omitempty"`

	// TopP implements nucleus sampling where only tokens with cumulative probability up to this value are considered.
	//
	// For example, 0.1 means only the tokens comprising the top 10% probability mass are considered.
	TopP float64 `json:"top_p,omitempty"`

	// Tools lists functions that the model can call, limited to 128.
	Tools []Tool `json:"tools,omitempty"`

	// ToolChoice specifies which tool to use if tools are provided.
	ToolChoice string `json:"tool_choice,omitempty"`

	// Logprobs indicates if log probabilities should be returned for each token.
	Logprobs bool `json:"logprobs,omitempty"`

	// TopLogprobs specifies how many most likely tokens to return with their log probabilities; requires Logprobs to be true.
	//
	// Range: 0 to 20.
	TopLogprobs float64 `json:"top_logprobs,omitempty"`
}

type Message struct {
	// Content is the message content.
	Content string `json:"content"`

	// Role is the role of the message author.
	Role Role `json:"role"`

	// Name is an optional name for the participant. It Provides the model information to differentiate between participants of the same role.
	Name string `json:"name,omitempty"`
}

type Role string

const (
	AssistantRole Role = "assistant"
	SystemRole    Role = "system"
	ToolRole      Role = "tool"
	UserRole      Role = "user"
)

// ResponseFormat specifies how the response should be formatted. By default the response is returned as plain text.
type ResponseFormat struct {
	// Type specifies the format of the response, either 'text' or 'json_object'.
	Type ResponseFormatType `json:"type,omitempty"`
}

type ResponseFormatType string

const (
	ResponseFormatText ResponseFormatType = "text"
	ResponseFormatJson ResponseFormatType = "json_object"
)

type Tool struct {
	// Type is the type of the tool.
	Type ToolType `json:"type"`

	// Function is the function details.
	Function ToolFunction `json:"function"`
}

type ToolType string

const ToolFunctionType ToolType = "function"

type ToolFunction struct {
	// Description provides a brief explanation of what the function does, aiding in its selection by the model.
	Description string `json:"description"`

	// Name is the name of the function.
	Name string `json:"name"`
}
type StreamOptions struct {
	IncludeUsage bool `json:"include_usage,omitempty"`
}

type CompletionResponse struct {
	// ID is a unique identifier for the chat completion.
	ID string `json:"id"`

	// Model indicates which model was used for this response.
	Model ModelID `json:"model"`

	// Choices contains one or more possible responses from the model.
	Choices []CompletionChoice `json:"choices"`

	// Created is the Unix timestamp (in seconds) of when the response was generated.
	Created int64 `json:"created"`

	// SystemFingerprint represents the backend configuration that the model runs with.
	SystemFingerprint string `json:"system_fingerprint"`

	// Usage is the usage statistics for the completion request.
	Usage CompletionUsage `json:"usage"`

	// Object describes the type of this response object i.e. "chat.completion" for a simple completion and "chat.completion.chunk" for a streaming completion.
	Object string `json:"object"`
}

type StreamCompletionChunk struct {
	// ID is a unique identifier for the chat completion.
	ID string `json:"id"`

	// Model indicates which model was used for this response.
	Model ModelID `json:"model"`

	// Choices contains one or more possible responses from the model.
	Choices []StreamCompletionChoice `json:"choices"`

	// Created is the Unix timestamp (in seconds) of when the response was generated.
	Created int64 `json:"created"`

	// SystemFingerprint represents the backend configuration that the model runs with.
	SystemFingerprint string `json:"system_fingerprint"`

	// Object describes the type of this response object i.e. "chat.completion" for a simple completion and "chat.completion.chunk" for a streaming completion.
	Object string `json:"object"`
}

type StreamCompletionChoice struct {
	Index        int64                  `json:"index"`
	Delta        StreamDelta            `json:"delta"`
	FinishReason CompletionFinishReason `json:"finish_reason"`
}

type StreamDelta struct {
	Content          string `json:"content"`
	ReasoningContent string `json:"reasoning_content"`
	Role             Role   `json:"role"`
}

type CompletionChoice struct {
	Index        int64                  `json:"index"`
	Message      Message                `json:"message"`
	FinishReason CompletionFinishReason `json:"finish_reason"`
}

type CompletionMessage struct {
	Content          string               `json:"content"`
	ReasoningContent string               `json:"reasoning_content"`
	Role             string               `json:"role"`
	ToolCalls        []CompletionToolCall `json:"tool_calls"`
}

type CompletionToolCall struct {
	ID       string                     `json:"id"`
	Type     ToolType                   `json:"type"`
	Function CompletionToolCallFunction `json:"function"`
}

type CompletionToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type CompletionFinishReason string

const (
	CompletionFinishReasonStop                       CompletionFinishReason = "stop"
	CompletionFinishReasonLength                     CompletionFinishReason = "length"
	CompletionFinishReasonContentFilter              CompletionFinishReason = "content_filter"
	CompletionFinishReasonToolCalls                  CompletionFinishReason = "tool_calls"
	CompletionFinishReasonInsufficientSystemResource CompletionFinishReason = "insufficient_system_resource"
)

type CompletionUsage struct {
	CompletionTokens        int64                       `json:"completion_tokens"`
	PromptTokens            int64                       `json:"prompt_tokens"`
	PromptCacheHitTokens    int64                       `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens   int64                       `json:"prompt_cache_miss_tokens"`
	TotalTokens             int64                       `json:"total_tokens"`
	CompletionTokensDetails CompletionUsageTokenDetails `json:"completion_tokens_details"`
}

type CompletionUsageTokenDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens"`
}
