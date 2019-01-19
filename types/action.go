package types

type Action struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Key    string `json:"key"`
	Data   string `json:"data"`
}
