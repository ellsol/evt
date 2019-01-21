package main

import (
	"fmt"
	"github.com/ellsol/evt/evt"
	"github.com/ellsol/evt/evtconfig"
	"github.com/ellsol/evt/utils"
	"log"
)

const testNetwork = "http://testnet1.everitoken.io:8888"
const mainNetwork = "https://mainnet2.everitoken.io"

const exampleTransactionId = "78b819cd91cb1ce81785503749f2cba983e5ae9551fe6cd1c9b077d53267d7a8"
const exampleBlockNum = "29042966"
const exampleBlockId = "01bb2916cc86cdd7dc8e4fe4e84115969d721ea2449a9fc5fe8c0f595c826283"


const exampleGroup = "vastchain.technology"
func main() {
	evt := evt.New(evtconfig.New(mainNetwork))
	//evt.Log.SetLevel(logrus.InfoLevel)


	PrintDomain("cookie", evt)
	//SearchForBlockWithTransaction(29000000, evt)

	//PrintGetInfo(evt)
	//PrintGetHeadBlockHeaderState(evt)

	//PrintGetTransactionIdsForBlock(evt, exampleBlockId)
	//PrintGetTransactionIdsForBlock(evt, "230")
	//PrintTransaction("78b819cd91cb1ce81785503749f2cba983e5ae9551fe6cd1c9b077d53267d7a8", evt)

	//PrintTransactionActions(exampleTransactionId, evt)
}


/////////////////////////////////////////////////
//  History
/////////////////////////////////////////////////
func PrintDomain(name string, evt *evt.Instance) {
	domain, err :=  evt.Api.V1.Evt.GetDomain("USDstable")


	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v\n",domain)
}

func PrintGroup(name string, evt *evt.Instance) {
	domain, err :=  evt.Api.V1.Evt.GetDomain("USDstable")


	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("%+v\n",domain)
}

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

/////////////////////////////////////////////////
//  History
/////////////////////////////////////////////////

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
