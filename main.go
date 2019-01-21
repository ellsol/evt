package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ellsol/evt/evt"
	"github.com/ellsol/evt/evtconfig"
	"github.com/ellsol/evt/utils"
	"github.com/sirupsen/logrus"
	"log"
)

const testNetwork = "http://testnet1.everitoken.io:8888"
const mainNetwork = "https://mainnet2.everitoken.io"

const exampleTransactionId = "78b819cd91cb1ce81785503749f2cba983e5ae9551fe6cd1c9b077d53267d7a8"
const exampleBlockNum = "29042966"
const exampleBlockId = "01bb2916cc86cdd7dc8e4fe4e84115969d721ea2449a9fc5fe8c0f595c826283"

const exampleGroup = "vastchain.technology"

func main() {
	spew.Config = spew.ConfigState{
		DisableCapacities:       false,
		DisablePointerMethods:   true,
		DisablePointerAddresses: true,
		SpewKeys:                false,
		Indent:                  "\t",
	}
	evt := evt.New(evtconfig.New(mainNetwork))
	evt.Log.SetLevel(logrus.TraceLevel)

	//PrintgetSuspend("suspend3", evt) // doesnt work

	//PrintFungibleBalance("EVT8MGU4aKiVzqMtWi9zLpu8KuTHZWjQQrX475ycSxEkLd6aBpraX", evt)
	//PrintFungible("3", evt)
	//PrintGetTokensByDomain("WALLETTEST", evt)
	//PrintDomain("cookie", evt)
	//SearchForBlockWithTransaction(29000000, evt)

	//PrintGetInfo(evt)
	//PrintGetHeadBlockHeaderState(evt)
	//PrintGetBlockHeaderState(evt, "01c10e6444fb14e270895b0bea2a148b10ff5c566732f858efadda70e8bac564")
	PrintGetTrxIdForLinkId("16951b9b653947955faa6c3cb3e506b6", evt)

	//PrintGetTransactionIdsForBlock(evt, exampleBlockId)
	//PrintGetTransactionIdsForBlock(evt, "230")
	//PrintTransaction("78b819cd91cb1ce81785503749f2cba983e5ae9551fe6cd1c9b077d53267d7a8", evt)

	//PrintTransactionActions(exampleTransactionId, evt)
}

/////////////////////////////////////////////////
//  Evt
/////////////////////////////////////////////////
func PrintgetSuspend(id string, evt *evt.Instance) {
	suspend, err := evt.Api.V1.Evt.GetSuspend(id)

	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(suspend)
}

func PrintFungibleBalance(address string, evt *evt.Instance) {
	fungibleBalance, err := evt.Api.V1.Evt.GetFungibleBalance(address)

	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(fungibleBalance)
}

func PrintFungible(id string, evt *evt.Instance) {
	fungible, err := evt.Api.V1.Evt.GetFungible(id)

	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(fungible)

	//utils.PrettyPrintStruct(fungible)
	//log.Printf("%+v\n", fungible)
}
func PrintGetToken(name string, evt *evt.Instance) {
	domain, err := evt.Api.V1.Evt.GetToken(name, name)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v\n", domain)
}

func PrintGetTokensByDomain(name string, evt *evt.Instance) {
	domain, err := evt.Api.V1.Evt.GetTokens(name, 0, 10)

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v\n", domain)
}

func PrintDomain(name string, evt *evt.Instance) {
	domain, err := evt.Api.V1.Evt.GetDomain("USDstable")

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v\n", domain)
}

func PrintGroup(name string, evt *evt.Instance) {
	domain, err := evt.Api.V1.Evt.GetDomain("USDstable")

	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v\n", domain)
}

/////////////////////////////////////////////////
//  History
/////////////////////////////////////////////////

func SearchForBlockWithTransaction(startBlock int, evt *evt.Instance) {

	for i := 0; i < 1000; i ++ {
		block, err := evt.Api.V1.Chain.GetBlock(fmt.Sprintf("%v", startBlock+i))
		if err != nil {
			log.Println(err)
		}

		if len(block.Transactions) > 0 {
			PrintGetBlock(evt, block.ID)
		}

	}
}

func PrintTransaction(transactionId string, evt *evt.Instance) {
	result, err := evt.Api.V1.History.GetTransaction(transactionId)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
}

func PrintTransactionActions(transactionId string, evt *evt.Instance) {
	result, err := evt.Api.V1.History.GetTransactionActions(transactionId)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
}

/////////////////////////////////////////////////
//  Chain
/////////////////////////////////////////////////
func PrintGetInfo(evt *evt.Instance) {
	info, err := evt.Api.V1.Chain.GetInfo()

	if err != nil {
		log.Println(err)
		return
	}

	utils.PrintColoredln("Version", info.EvtAPIVersion)
	utils.PrintColoredln("ChainId", info.ChainID)
	utils.PrintColoredln("Block Id", info.HeadBlockID)
	utils.PrintColoredln("Block Num", info.HeadBlockNum)
	utils.PrintColoredln("Block Producer", info.HeadBlockProducer)
	utils.PrintColoredln("block time", info.HeadBlockTime)
}

func PrintGetHeadBlockHeaderState(evt *evt.Instance) {
	headerState, err := evt.Api.V1.Chain.GetHeadBlockHeaderState()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(headerState.BlockNum)
}

func PrintGetBlockHeaderState(evt *evt.Instance, blockNumOrId string) {
	result, err := evt.Api.V1.Chain.GetBlockHeaderState(blockNumOrId)

	if err != nil {
		log.Println(err)
		return
	}
	spew.Dump(result)
}

func PrintGetBlock(evt *evt.Instance, blockNumOrId string) {
	result, err := evt.Api.V1.Chain.GetBlock(blockNumOrId)

	if err != nil {
		log.Println(err)
		return
	}

	utils.PrintColoredln("BlockNum", result.BlockNum)
	utils.PrintColoredln("BlockId", result.ID)
	utils.PrintColoredln("Timestamp", result.Timestamp)
	utils.PrintColoredln("Transactions", "")
	for _, t := range result.Transactions {
		utils.PrintColoredln("Id", t.Trx.ID)
		utils.PrintColoredln("Type", t.Type)
		utils.PrintColoredln("Status", t.Status)
	}
	//spew.Dump(result)
}

func PrintGetTransactionIdsForBlock(evt *evt.Instance, blockId string) {
	result, err := evt.Api.V1.Chain.GetTransactionIdsForBlock(blockId)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
}

func PrintGetTrxIdForLinkId(linkId string, evt *evt.Instance) {
	result, err := evt.Api.V1.Chain.GetTrxIdForLinkId(linkId)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
}

