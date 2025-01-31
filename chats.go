package deepseek

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/roushou/deepseek/internal/http_client"
)

type ChatsClient struct {
	httpClient *http_client.Client
}

func (c *ChatsClient) CreateCompletion(args ChatCompletionArgs) (*ChatCompletionResponse, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	req, err := c.httpClient.NewRequest(http.MethodPost, "/chat/completions", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	var completion ChatCompletionResponse
	err = c.httpClient.Do(req, &completion)
	return &completion, err
}

// NewChatCompletionRequest creates a new chat completion request with default values.
func NewChatCompletionRequest(model ModelID) ChatCompletionArgs {
	return ChatCompletionArgs{
		Model:            model,
		Messages:         []Message{},
		FrequencyPenalty: 0,
		MaxTokens:        4096,
		PresencePenalty:  0,
		ResponseFormat:   &ResponseFormat{},
		Stop:             []string{},
		Stream:           false,
		StreamOptions:    &StreamOptions{},
		Temperature:      1,
		TopP:             1,
		Tools:            []Tool{},
		ToolChoice:       "",
		Logprobs:         false,
		TopLogprobs:      0,
	}
}

type ChatCompletionArgs struct {
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

type ChatCompletionResponse struct {
	// ID is a unique identifier for the chat completion.
	ID string `json:"id"`

	// Model indicates which model was used for this response.
	Model ModelID `json:"model"`

	// Choices contains one or more possible responses from the model.
	Choices []ChatCompletionChoice `json:"choices"`

	// Created is the Unix timestamp (in seconds) of when the response was generated.
	Created int64 `json:"created"`

	// SystemFingerprint represents the backend configuration that the model runs with.
	SystemFingerprint string `json:"system_fingerprint"`

	// Usage is the usage statistics for the completion request.
	Usage ChatCompletionUsage `json:"usage"`

	// Object describes the type of this response object. Always "chat.completion" in this case.
	Object string `json:"object"`
}

type ChatCompletionChoice struct {
	Index        int64                      `json:"index"`
	Message      Message                    `json:"message"`
	FinishReason ChatCompletionFinishReason `json:"finish_reason"`
}

type ChatCompletionMessage struct {
	Content          string                   `json:"content"`
	ReasoningContent string                   `json:"reasoning_content"`
	Role             string                   `json:"role"`
	ToolCalls        []ChatCompletionToolCall `json:"tool_calls"`
}

type ChatCompletionToolCall struct {
	ID       string                         `json:"id"`
	Type     ToolType                       `json:"type"`
	Function ChatCompletionToolCallFunction `json:"function"`
}

type ChatCompletionToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ChatCompletionFinishReason string

const (
	CompletionFinishReasonStop                       ChatCompletionFinishReason = "stop"
	CompletionFinishReasonLength                     ChatCompletionFinishReason = "length"
	CompletionFinishReasonContentFilter              ChatCompletionFinishReason = "content_filter"
	CompletionFinishReasonToolCalls                  ChatCompletionFinishReason = "tool_calls"
	CompletionFinishReasonInsufficientSystemResource ChatCompletionFinishReason = "insufficient_system_resource"
)

type ChatCompletionUsage struct {
	CompletionTokens        int64                           `json:"completion_tokens"`
	PromptTokens            int64                           `json:"prompt_tokens"`
	PromptCacheHitTokens    int64                           `json:"prompt_cache_hit_tokens"`
	PromptCacheMissTokens   int64                           `json:"prompt_cache_miss_tokens"`
	TotalTokens             int64                           `json:"total_tokens"`
	CompletionTokensDetails ChatCompletionUsageTokenDetails `json:"completion_tokens_details"`
}

type ChatCompletionUsageTokenDetails struct {
	ReasoningTokens int64 `json:"reasoning_tokens"`
}
