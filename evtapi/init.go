package evtapi

import (
	"github.com/ellsol/evt/evtapi/client"
	"github.com/ellsol/evt/evtapi/v1"
	"github.com/ellsol/evt/evtconfig"
	"github.com/sirupsen/logrus"
)

type Instance struct {
	V1     *v1.Instance
}

func New(config *evtconfig.Instance, logger *logrus.Logger) *Instance {
	c := client.New(config, logger)

	return &Instance{
		V1: v1.New(config, c),
	}
}
