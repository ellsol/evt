package transaction

import "github.com/ellsol/evt/evttypes"

const actionNameTransferFungible = "transferft"
const fungibleDomain = ".fungible"

type TransferFungible struct {
	from       string
	to         string
	asset      *evttypes.Asset
	fungibleId string
	memo       string
}

func NewTransferFungible(from string, to string, value string, fungibleId string) *TransferFungible {
	return &TransferFungible{
		from,
		to,
		evttypes.NewAsset(value, fungibleId),
		fungibleId,
		"",
	}
}

func (it *TransferFungible) SetMemo(memo string) *TransferFungible {
	it.memo = memo
	return it
}

type TransferFungibleArguments struct {
	From   string `json:"from"`   // address
	To     string `json:"to"`     // address
	Number string `json:"number"` // asset
	Memo   string `json:"memo"`   // string
}

func (it *TransferFungible) Arguments() *evttypes.ActionArguments {
	arg := TransferFungibleArguments{
		From:   it.from,
		To:     it.to,
		Number: it.asset.String(),
		Memo:   it.memo,
	}

	return &evttypes.ActionArguments{
		Action: actionNameTransferFungible,
		Args:   arg,
	}
}

func (it *TransferFungible) Action(binargs string) *evttypes.SimpleAction {
	return &evttypes.SimpleAction{
		Data: binargs,
		Action: evttypes.Action{
			Name:   actionNameTransferFungible,
			Domain: fungibleDomain,
			Key:    it.fungibleId,
		},
	}
}
