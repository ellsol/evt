package ecc

import (
	"io/ioutil"
)

func (it *PrivateKey) Save(filename string) error {
	err := ioutil.WriteFile(filename, []byte(it.String()), 0644)

	if err != nil {
		return err
	}

	return nil
}

func LoadPrivateKey(filename string) (*PrivateKey, error) {
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return NewPrivateKey(string(b))
}
