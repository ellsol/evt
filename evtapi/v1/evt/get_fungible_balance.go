package evt

import (
	"github.com/ellsol/evt/evtapi/client"
)

type GetFungibleBalanceRequest struct {
	Address string `json:"address"`
}

type GetFungibleBalanceResult = []string

func (it *Instance) GetFungibleBalance(id string) (*GetFungibleBalanceResult, *client.ApiError) {
	response := &GetFungibleBalanceResult{}

	err := it.client.Post(it.path("get_fungible_balance"), &GetFungibleBalanceRequest{id}, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
