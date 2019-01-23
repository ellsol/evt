package transaction

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/ellsol/evt/ecc"
	"github.com/ellsol/evt/evt"
	"github.com/ellsol/evt/evtapi/v1/chain"
	"github.com/ellsol/evt/evttypes"
	"github.com/ellsol/evt/utils"
	"github.com/sirupsen/logrus"
)

func Deploy(action EvtAction, privateKey *ecc.PrivateKey, evt *evt.Instance) (*chain.PushTransactionResult, error) {
	abiJsonToBinResult, apiError := evt.Api.V1.Chain.AbiJsonToBin(action.Arguments())

	if apiError != nil {
		logrus.Error(apiError)
		return nil, apiError.Error()
	}

	trxJson, err := getTrxJsonBase(privateKey, evt)

	if err != nil {
		return nil, err
	}

	trxJson.Actions = []evttypes.SimpleAction{*action.Action(abiJsonToBinResult.Binargs)}

	return signAndPostTransaction(trxJson, privateKey, evt)
}

func getTrxJsonBase(privateKey *ecc.PrivateKey, evt *evt.Instance) (*evttypes.TRXJson, error) {
	info, apiError := evt.Api.V1.Chain.GetInfo()

	if apiError != nil {
		return nil, apiError.Error()
	}

	refBlockNum, refBlockPrefix := getNumAndRefFromBlockID(info.LastIrreversibleBlockID)

	return &evttypes.TRXJson{
		RefBlockNum:           int(refBlockNum),
		RefBlockPrefix:        int(refBlockPrefix),
		Payer:                 privateKey.PublicKey().String(),
		Expiration:            utils.In5Mins(),
		MaxCharge:             10000,
		TransactionExtensions: make([]interface{}, 0),
	}, nil
}

func signAndPostTransaction(trxJson *evttypes.TRXJson, privKey *ecc.PrivateKey, evt *evt.Instance) (*chain.PushTransactionResult, error) {
	// Step 1 Get Transaction Digest
	digest, apiError := evt.Api.V1.Chain.TRXJsonToDigest(trxJson)

	if apiError != nil {
		logrus.Error(apiError)
		return nil, apiError.Error()
	}

	logrus.Tracef("Received Digest: %v\n", digest.Digest)

	b, err := hex.DecodeString(digest.Digest)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	// Step 2 Sign Transaction
	signature, err := privKey.Sign(b)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logrus.Tracef("Signed Transaction: ", signature.String())

	// Step 3 Push Transaction
	pushTransactionResult, apiError := evt.Api.V1.Chain.PushTransaction([]string{signature.String()}, trxJson)

	if apiError != nil {
		logrus.Println(apiError.String())
		return nil, apiError.Error()
	}

	logrus.Tracef("Transaction successfully posted: ", pushTransactionResult.TransactionId)

	return pushTransactionResult, nil
}


func getNumAndRefFromBlockID(lastReversibleblockId string) (int, int) {
	headBlockId, err := hex.DecodeString(lastReversibleblockId)

	if err != nil {
		logrus.Println(err)
		return -1, -1
	}

	refBlockNum := binary.BigEndian.Uint16(headBlockId[2:4])
	refBlockPrefix := binary.LittleEndian.Uint32(headBlockId[8:])

	return int(refBlockNum), int(refBlockPrefix)
}

