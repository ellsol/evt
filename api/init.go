package api

import (
	"github.com/ellsol/evt/api/client"
	"github.com/ellsol/evt/api/v1"
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
