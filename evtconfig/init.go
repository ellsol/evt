package evtconfig

import (
	"github.com/sirupsen/logrus"
)

const version1 = "v1"

type Instance struct {
	HttpPath string
	Version  string

	Log *logrus.Logger
}

func New(httpPath string) *Instance {
	return &Instance{
		HttpPath: httpPath,
		Version:  version1,
		Log:      logrus.New(),
	}
}
