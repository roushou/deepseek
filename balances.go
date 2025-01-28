package deepseek

import (
	"net/http"

	"github.com/roushou/deepseek/internal/http_client"
)

type BalancesClient struct {
	httpClient *http_client.Client
}

func (c *BalancesClient) GetUserBalance() (*UserBalanceResponse, error) {
	req, err := c.httpClient.NewRequest(http.MethodGet, "/user/balance", nil)
	if err != nil {
		return nil, err
	}
	var balance UserBalanceResponse
	err = c.httpClient.Do(req, &balance)
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

type UserBalanceResponse struct {
	IsAvailable  bool              `json:"is_available"`
	BalanceInfos []UserBalanceInfo `json:"balance_infos"`
}

type UserBalanceInfo struct {
	Currency        string `json:"currency"`
	TotalBalance    string `json:"total_balance"`
	GrantedBalance  string `json:"granted_balance"`
	ToppedUpBalance string `json:"topped_up_balance"`
}
