package chain

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evttypes"
)

type PushTransactionResult struct {
	TransactionId string `json:"transaction_id"`
}

func (it *Instance) PushTransaction(signedTRXJson *evttypes.SignedTRXJson) (*PushTransactionResult, *client.ApiError) {
	result := &PushTransactionResult{}

	err := it.Client.Post(it.Path("push_transaction"), signedTRXJson, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
