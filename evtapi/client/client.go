package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ellsol/evt/evtconfig"
	"io/ioutil"
	"log"
	"net/http"
)

type Instance struct {
	config *evtconfig.Instance
}

func New(config *evtconfig.Instance) *Instance {
	return &Instance{
		config: config,
	}
}

func (it *Instance) Post(path string, body interface{}, response interface{}) *ApiError {
	bbody, err := json.Marshal(body)

	if err != nil {
		return NewApiError(fmt.Errorf("http post parsing error: %v", err))
	}

	resp, err := http.Post(it.getUrl(path), "application/json", bytes.NewReader(bbody))
	defer resp.Body.Close()

	if err != nil {
		return NewApiError(fmt.Errorf("http post: %v", err))
	}

	b, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode >= 300 {
		return parseError(b)
	}

	log.Println("Raw: ", string(b))

	if err != nil {
		return NewApiError(fmt.Errorf("http post: %v", err))
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

	it.config.Log.Printf("http get request url: %v\n", url)

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
