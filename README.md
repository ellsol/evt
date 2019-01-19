# everiToken API library for Golang

## Includes 

1. everitoken Api Client 
2. everiToken Wallet
3. Processes for posting actions

## Basic Usage

Create a configuration and the sdk

    config := evtconfig.New(httpPath)
    sdk := evtsdk.New(config)
    result, err := sdk.api.v1.chain.GetBlock()
 
    
