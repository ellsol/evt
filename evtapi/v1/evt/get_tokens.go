package evt

import (
	"github.com/ellsol/evt/evtapi/client"
)

type GetTokensRequest struct {
	Name string `json:"name"`
}

type GetTokensResult struct {
}

func (it *Instance) GetTokens(domainName string) (*GetTokensResult, *client.ApiError) {
	response := &GetTokensResult{}

	err := it.client.Post(it.path("get_tokens"), &GetTokensRequest{domainName}, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
