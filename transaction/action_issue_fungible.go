package transaction

import "github.com/ellsol/evt/evttypes"

const actionNameIssueFungible = "issuefungible"

type IssueFungible struct {
	address    string
	asset      *evttypes.Asset
	fungibleId string
	memo       string
}

func NewIssueFungible(address string, value string, fungibleId string) *IssueFungible {
	return &IssueFungible{
		address,
		evttypes.NewAsset(value, fungibleId),
		fungibleId,
		"",
	}
}

func (it *IssueFungible) SetMemo(memo string) *IssueFungible {
	it.memo = memo
	return it
}

type IssueFungibleArguments struct {
	Address string `json:"address"` // address
	Number  string `json:"number"`  // asset
	Memo    string `json:"memo"`    // string
}

func (it *IssueFungible) Arguments() *evttypes.ActionArguments {
	arg := IssueFungibleArguments{
		Address: it.address,
		Number:  it.asset.String(),
		Memo:    it.memo,
	}

	return &evttypes.ActionArguments{
		Action: actionNameIssueFungible,
		Args:   arg,
	}
}

func (it *IssueFungible) Action(binargs string) *evttypes.SimpleAction {
	return &evttypes.SimpleAction{
		Data: binargs,
		Action: evttypes.Action{
			Name:   actionNameIssueFungible,
			Domain: fungibleDomain,
			Key:    it.fungibleId,
		},
	}
}
