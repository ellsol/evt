# Golang SDK for everiToken

## Includes 

1. everitoken Api Client 
2. everiToken Wallet
3. Processes for posting actions

## Init

Create a configuration and the sdk

    config := evtconfig.New(httpPath)
    sdk := sdk.New(config)

### Api

e.g.

    result, err := sdk.api.v1.chain.GetBlock()
    
