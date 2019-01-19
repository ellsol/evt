package chain

import (
	"fmt"
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evtconfig"
)

type Instance struct {
	Client *client.Instance
	Config *evtconfig.Instance
}

func New(config *evtconfig.Instance, client *client.Instance) *Instance {
	return &Instance{
		Client: client,
		Config: config,
	}
}

func (it *Instance) Path(method string) string {
	return fmt.Sprintf("chain/%v", method)
}
