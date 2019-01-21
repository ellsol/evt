package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ellsol/evt/evtconfig"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type Instance struct {
	config *evtconfig.Instance
	logger *logrus.Logger
}

func New(config *evtconfig.Instance, logger *logrus.Logger) *Instance {
	return &Instance{
		config: config,
		logger: logger,
	}
}

func (it *Instance) Post(path string, body interface{}, response interface{}) *ApiError {
	url := it.getUrl(path)
	it.logger.Infof("client: posting to %v with body %+v\n", url, body)

	bbody, err := json.Marshal(body)

	if err != nil {
		return NewApiError(fmt.Errorf("client:Post parsing error %v", err))
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(bbody))

	if err != nil {
		return NewApiError(fmt.Errorf("client:Post request error %v", err))
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return parseError(b)
	}

	//log.Println("Raw: ", string(b))

	if err != nil {
		return NewApiError(fmt.Errorf("client:Post parsing response error %v", err))
	}

	err = json.Unmarshal(b, &response)

	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (it *Instance) Get(path string, response interface{}) *ApiError {
	url := it.getUrl(path)
	resp, err := http.Get(url)

	it.logger.Infof("client: get %v\n", url)

	if err != nil {
		return NewApiError(fmt.Errorf("http get request: %v", err))
	}

	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		return parseError(b)
	}

	if err != nil {
		return NewApiError(fmt.Errorf("http get: %v", err))
	}

	err = json.Unmarshal(b, &response)

	if err != nil {
		return NewApiError(err)
	}

	return nil
}

func (it *Instance) getUrl(path string) string {
	return fmt.Sprintf("%v/%v/%v", it.config.HttpPath, it.config.Version, path)
}
