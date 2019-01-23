package evttypes

import "fmt"

type Role struct {
	Name        string       `json:"name"`
	Threshold   int          `json:"threshold"`
	Authorizers []Authorizer `json:"authorizers"`
}

type Authorizer struct {
	Ref    string `json:"ref"`
	Weight int    `json:"weight"`
}

func GroupOwnedAuthorizer() *Authorizer {
	return &Authorizer{
		Ref:    "[G] Owner",
		Weight: 1,
	}
}

func SingleAddressAuthorizer(address string) *Authorizer {
	return &Authorizer{
		Ref:    fmt.Sprintf("[A] %v", address),
		Weight: 1,
	}
}
