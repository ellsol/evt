package chain

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evttypes"
)

type AbiJsonToBinRequest = evttypes.ActionArguments

type AbiJsonToBinResult struct {
	Binargs string `json:"binargs"`
}

type Args struct {
	Name     string     `json:"name"`
	Creator  string     `json:"creator"`
	Issue    ActionType `json:"issue"`
	Transfer ActionType `json:"transfer"`
	Manage   ActionType `json:"manage"`
}

type Authorizers struct {
	Ref    string `json:"ref"`
	Weight int    `json:"weight"`
}

type ActionType struct {
	Name        string        `json:"name"`
	Threshold   int           `json:"threshold"`
	Authorizers []Authorizers `json:"authorizers"`
}

func (it *Instance) AbiJsonToBin(request *AbiJsonToBinRequest) (*AbiJsonToBinResult, *client.ApiError) {
	response := &AbiJsonToBinResult{}

	err := it.Client.Post(it.Path("abi_json_to_bin"), request, response)

	if err != nil {
		return nil, err
	}

	return response, nil
}
