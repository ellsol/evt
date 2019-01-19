package chain

import "github.com/ellsol/evt/api/client"

type GetBlockRequest struct {
	BlockNumOrID string `json:"block_num_or_id"`
}

type GetBlockResult struct {
	Timestamp         string        `json:"timestamp"`
	Producer          string        `json:"producer"`
	Confirmed         int           `json:"confirmed"`
	Previous          string        `json:"previous"`
	TransactionMroot  string        `json:"transaction_mroot"`
	ActionMroot       string        `json:"action_mroot"`
	ScheduleVersion   int           `json:"schedule_version"`
	NewProducers      interface{}   `json:"new_producers"`
	HeaderExtensions  []interface{} `json:"header_extensions"`
	ProducerSignature string        `json:"producer_signature"`
	Transactions      []interface{} `json:"transactions"`
	BlockExtensions   []interface{} `json:"block_extensions"`
	ID                string        `json:"id"`
	BlockNum          int           `json:"block_num"`
	RefBlockPrefix    int           `json:"ref_block_prefix"`
}

func (it *Instance) GetBlock(request *GetBlockRequest) (*GetBlockResult, *client.ApiError) {
	response := &GetBlockResult{}

	err := it.Client.Post(it.Path("get_block"), request, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
