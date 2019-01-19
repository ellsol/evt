package evtapi

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evtapi/v1"
	"github.com/ellsol/evt/evtconfig"
)

type Instance struct {
	Client *client.Instance
	V1     *v1.Instance
}

func New(config *evtconfig.Instance) *Instance {
	c := client.New(config)

	return &Instance{
		V1: v1.New(config, c),
	}
}
