package transaction

import (
	"github.com/ellsol/evt/ecc"
	"github.com/ellsol/evt/evttypes"
)

type TransactionMaker struct {
	Name     string
	Creator  *ecc.PrivateKey
	Issue    *evttypes.Role
	Transfer *evttypes.Role
	Manage   *evttypes.Role
}

func CreateTransaction(name string, creator *ecc.PrivateKey) *TransactionMaker {
	return &TransactionMaker{
		Name:    name,
		Creator: creator,
	}
}

func (it *TransactionMaker) SetIssue(treshold int, authorizer *evttypes.Authorizer) *TransactionMaker {
	it.Issue = &evttypes.Role{
		Name:        "issue",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}

func (it *TransactionMaker) SetTransfer(treshold int, authorizer *evttypes.Authorizer) *TransactionMaker {
	it.Transfer = &evttypes.Role{
		Name:        "transfer",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}

func (it *TransactionMaker) SetManage(treshold int, authorizer *evttypes.Authorizer) *TransactionMaker {
	it.Manage = &evttypes.Role{
		Name:        "manage",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}


