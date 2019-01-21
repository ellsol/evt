package history

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evtapi/v1/chain"
)

type GetTransactionRequest struct {
	TransactionId string `json:"id"`
}

type GetTransactionResult = chain.Transaction

func (it *Instance) GetTransaction(transactionId string) (*GetTransactionActionsRequest, *client.ApiError) {
	response := &GetTransactionActionsRequest{}

	err := it.Client.Post(it.Path("get_transaction_actions"), &GetTransactionRequest{transactionId}, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
