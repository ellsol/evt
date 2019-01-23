package transaction

import (
	"github.com/ellsol/evt/ecc"
	"github.com/ellsol/evt/evttypes"
)

type CreateDomain struct {
	Name     string
	Creator  *ecc.PrivateKey
	Issue    *evttypes.Role
	Transfer *evttypes.Role
	Manage   *evttypes.Role
}

func NewCreateDomain(name string, creator *ecc.PrivateKey) *CreateDomain {
	return &CreateDomain{
		Name:    name,
		Creator: creator,
	}
}

func (it *CreateDomain) SetIssue(treshold int, authorizer *evttypes.Authorizer) *CreateDomain {
	it.Issue = &evttypes.Role{
		Name:        "issue",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}

func (it *CreateDomain) SetTransfer(treshold int, authorizer *evttypes.Authorizer) *CreateDomain {
	it.Transfer = &evttypes.Role{
		Name:        "transfer",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}

func (it *CreateDomain) SetManage(treshold int, authorizer *evttypes.Authorizer) *CreateDomain {
	it.Manage = &evttypes.Role{
		Name:        "manage",
		Threshold:   treshold,
		Authorizers: []evttypes.Authorizer{*authorizer},
	}
	return it
}


