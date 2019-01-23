package chain

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evttypes"
)

type TRXJsonToDigestRequest = evttypes.TRXJson

type TRXJsonToDigestResult struct {
	Digest string `json:"digest"`
}

func (it *Instance) TRXJsonToDigest(trxJson *evttypes.TRXJson) (*TRXJsonToDigestResult, *client.ApiError) {
	response := &TRXJsonToDigestResult{}

	err := it.Client.Post(it.Path("trx_json_to_digest"), trxJson, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
