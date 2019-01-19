package chain

import "github.com/ellsol/evt/api/client"

type GetBlockRequest struct {
	BlockNumOrID string `json:"block_num_or_id"`
}

type GetBlockResult struct  {

}

func (it *Instance) GetBlock(request *GetBlockRequest) (*GetBlockResult, *client.ApiError) {
	response := &GetBlockResult{}

	err := it.Client.Post(it.Path("get_block"), request, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}

