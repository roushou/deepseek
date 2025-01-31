package deepseek

import (
	"fmt"
	"net/http"

	"github.com/roushou/deepseek/internal/http_client"
)

type ModelID string

const (
	DeepSeekChat     ModelID = "deepseek-chat"
	DeepSeekCoder    ModelID = "deepseek-coder"
	DeepSeekReasoner ModelID = "deepseek-reasoner"
)

type ModelsClient struct {
	httpClient *http_client.Client
}

// Model struct represents a DeepSeek model.
type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}

// ModelsList is a list of models, including those that belong to the user or organization.
type ModelsList struct {
	Data []Model `json:"data"`
}

// ListModels Lists the currently available models and provides basic information about each model such as the model id and parent.
func (c *ModelsClient) ListModels() (*ModelsList, error) {
	var models ModelsList
	req, err := c.httpClient.NewRequest(http.MethodGet, "/models", nil)
	if err != nil {
		return nil, err
	}
	_, err = c.httpClient.Do(req, &models)
	return &models, err
}

// GetModel Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
func (c *ModelsClient) GetModel(modelID string) (*Model, error) {
	var model Model
	path := fmt.Sprintf("/models/%s", modelID)
	req, err := c.httpClient.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	_, err = c.httpClient.Do(req, &model)
	return &model, err
}
