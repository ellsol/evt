package v1

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evtapi/v1/chain"
	"github.com/ellsol/evt/evtconfig"
)

type Instance struct {
	Chain *chain.Instance
}

func New(config *evtconfig.Instance, instance *client.Instance) *Instance {
	return &Instance{
		Chain: chain.New(config, instance),
	}
}
