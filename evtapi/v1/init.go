package v1

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evtapi/v1/chain"
	"github.com/ellsol/evt/evtapi/v1/evt"
	"github.com/ellsol/evt/evtapi/v1/history"
	"github.com/ellsol/evt/evtconfig"
)

type Instance struct {
	Chain   *chain.Instance
	Evt     *evt.Instance
	History *history.Instance
}

func New(config *evtconfig.Instance, client *client.Instance) *Instance {
	return &Instance{
		Chain:   chain.New(config, client),
		Evt:     evt.New(config, client),
		History: history.New(config, client),
	}
}
