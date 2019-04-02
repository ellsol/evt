# Fungible

A fungible is an interchangable token. To issue a fungible it needs
to be created first and then issued.  


## Create Fungible

    manager := evttypes.SingleAddressAuthorizer(privateKey.PublicKey().String())
    issuer := evttypes.SingleAddressAuthorizer(privateKey.PublicKey().String())
    
	nf := fungible.New(
		"MyCoin",                          // name
		privateKey.PublicKey().String(),   // public key of creator
		"MC",                              // symbol name
		"9000",                            // fungible id
		4,                                 // precision
		"2000000.0000").                   // total supply
		SetManageRole(1, manager).         // manager of fungible (can change meta)
		SetIssueRole(1, issuer)            // issuer of fungible (can issue tokens)


	transactionResult, err := transaction.Deploy(nf, privKey, evt)
	

#### Issue Fungible


	ifu := fungible.Issue(
		privKey.PublicKey().String(),      // public key of issuer
		"2000000.0000",                    // amount to be issued
		"9000")                            // fungible id
	
	transactionResult, err := transaction.Deploy(ifu, privateKey, evt)

#### Transfer Fungible to someone
    
    
	tf := fungible.Transfer(
		"EVT6f4...",                        // from address
		"EVT134...",                        // to address
		"10.00",                            // amount
		"9000")                             // fungible id

	transactionResult, err := transaction.Deploy(tf, privateKey, evtinstance)
	