package api

import "github.com/ellsol/evt/evtconfig"

type Instance struct {
	Config *evtconfig.Instance
}

func New(config *evtconfig.Instance) *Instance {
	return &Instance{
		Config: config,
	}
}
