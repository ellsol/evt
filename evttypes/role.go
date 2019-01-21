package evttypes

type Role struct {
	Name        string       `json:"name"`
	Threshold   int          `json:"threshold"`
	Authorizers []Authorizer `json:"authorizers"`
}

type Authorizer struct {
	Ref    string `json:"ref"`
	Weight int    `json:"weight"`
}
