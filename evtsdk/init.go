package evtsdk

import (
	"github.com/ellsol/evt/evtapi"
	"github.com/ellsol/evt/evtconfig"
)

type Instance struct {
	Api *evtapi.Instance
}

func New(config *evtconfig.Instance) *Instance {
	return &Instance{
		Api: evtapi.New(config),
	}
}

