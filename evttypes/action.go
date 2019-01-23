package evttypes

type Action struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	Key    string `json:"key"`
}

type SimpleAction struct {
	Action
	Data string `json:"data"`
}

type FullAction struct {
	Action
	Data    Args   `json:"data"`
	HexData string `json:"hex_data"`
}

type ActionArguments struct {
	Action string      `json:"action"`
	Args   interface{} `json:"args"`
}

type Args struct {
	Name        string `json:"name"`
	Creator     string `json:"creator"`
	Issue       *Role  `json:"issue,omitempty"`
	Transfer    *Role  `json:"transfer,omitempty"`
	Manage      *Role  `json:"manage,omitempty"`
	TotalSupply int64  `json:"total_supply,omitempty"`
}
