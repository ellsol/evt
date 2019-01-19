package chain

import "github.com/ellsol/evt/api/client"

type GetTransactionIdsForBlockRequest struct {
	BlockID string `json:"block_id"`
}

type GetTransactionIdsForBlockResult = []string

func (it *Instance) GetTransactionIdsForBlock(request *GetTransactionIdsForBlockRequest) (*GetTransactionIdsForBlockResult, *client.ApiError) {
	response := &GetTransactionIdsForBlockResult{}

	err := it.Client.Post(it.Path("get_transaction_ids_for_block"), request, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
