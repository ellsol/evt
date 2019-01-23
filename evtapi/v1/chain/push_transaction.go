package chain

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evttypes"
)

type PushTransactionRequest struct {
	Signatures  []string          `json:"signatures"`
	Compression string            `json:"compression"`
	Transaction *evttypes.TRXJson `json:"transaction"`
}

type PushTransactionResult struct {
	TransactionId string `json:"transaction_id"`
}

func (it *Instance) PushTransaction(signatures []string, trxJson *evttypes.TRXJson) (*PushTransactionResult, *client.ApiError) {
	result := &PushTransactionResult{}

	err := it.Client.Post(it.Path("push_transaction"), &PushTransactionRequest{signatures, "none", trxJson}, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}
