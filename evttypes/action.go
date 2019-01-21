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
	Data struct {
		Name     string `json:"name"`
		Creator  string `json:"creator"`
		Issue    Role   `json:"issue"`
		Transfer Role   `json:"transfer"`
		Manage   Role   `json:"manage"`
	} `json:"data"`
	HexData string `json:"hex_data"`
}
