package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ellsol/evt/api/v1/chain"
	"github.com/ellsol/evt/evtconfig"
	"github.com/ellsol/evt/sdk"
	"log"
)

const testNetwork = "http://testnet1.everitoken.io:8888"
const mainNetwork = "http://mainnet2.everitoken.io"

func main() {
	evt := sdk.New(evtconfig.New(testNetwork))


	info, err := evt.Api.V1.Chain.GetInfo()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(info.EvtAPIVersion)

	headerState, err := evt.Api.V1.Chain.GetHeadBlockHeaderState()

	if err != nil {
		log.Println(err)
		return
	}

	log.Println(headerState.BlockNum)

	result, err := evt.Api.V1.Chain.GetBlock(&chain.GetBlockRequest{
		BlockNumOrID: fmt.Sprintf("%v", headerState.BlockNum),
	})

	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(result)
}
