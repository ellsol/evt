package chain

import "github.com/ellsol/evt/evtapi/client"

type GetTransactionRequest struct {
	BlockNum string `json:"block_num"`
	Id       string `json:"id"`
}

type GetTransactionResult struct {
}

func (it *Instance) GetTransaction(blockNum string, id string) (*GetTransactionResult, *client.ApiError) {
	response := &GetTransactionResult{}

	err := it.Client.Get(it.Path("get_info"), response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
