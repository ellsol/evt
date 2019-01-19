package sdk

import (
	"github.com/ellsol/evt/api"
	"github.com/ellsol/evt/evtconfig"
)

type Instance struct {
	Api *api.Instance
}

func New(config *evtconfig.Instance) *Instance {
	return &Instance{
		Api:api.New(config),
	}
}

