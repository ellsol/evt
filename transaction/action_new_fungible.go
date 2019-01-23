package transaction

import "github.com/ellsol/evt/evttypes"

const actionNameNewFungible = "newfungible"

type NewFungible struct {
	Name        string
	Creator     string
	SymName     string
	Symbol      *evttypes.Symbol
	FungibleId  string
	TotalSupply *evttypes.Asset
	Issue       *evttypes.Role
	Manage      *evttypes.Role
}

func NewNewFungible(name string, creator string, symName string, fungibleId string, precision int, supply string) *NewFungible {
	return &NewFungible{
		Name:        name,
		Creator:     creator,
		SymName:     symName,
		FungibleId:  fungibleId,
		Symbol:      evttypes.NewSymbol(precision, fungibleId),
		TotalSupply: evttypes.NewAsset(supply, fungibleId),
	}
}

func (it *NewFungible) SetManageRole(treshold int, authorizer *evttypes.Authorizer) *NewFungible {
	it.Manage = &evttypes.Role{
		Name:        "manage",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}

func (it *NewFungible) SetIssueRole(treshold int, authorizer *evttypes.Authorizer) *NewFungible {
	it.Issue = &evttypes.Role{
		Name:        "issue",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}

type NewFungibleArguments struct {
	Name        string         `json:"name"`         // fungible_name
	SymName     string         `json:"sym_name"`     // symbol_name
	Sym         string         `json:"sym"`          // symbol
	Creator     string         `json:"creator"`      // user_id
	Issue       *evttypes.Role `json:"issue"`        // permission_def
	Manage      *evttypes.Role `json:"manage"`       // permission_def
	TotalSupply string         `json:"total_supply"` // asset
}

// EvtAction Implementation

func (it *NewFungible) Arguments() *evttypes.ActionArguments {
	return &evttypes.ActionArguments{
		Action: actionNameNewFungible,
		Args: NewFungibleArguments{
			Name:        it.Name,
			Creator:     it.Creator,
			Issue:       it.Issue,
			Manage:      it.Manage,
			Sym:         it.Symbol.String(),
			TotalSupply: it.TotalSupply.String(),
			SymName:     it.SymName,
		},
	}
}

func (it *NewFungible) Action(binargs string) *evttypes.SimpleAction {
	return &evttypes.SimpleAction{
		Data: binargs,
		Action: evttypes.Action{
			Name:   actionNameNewFungible,
			Domain: fungibleDomain,
			Key:    it.FungibleId,
		},
	}
}
