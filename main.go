package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"log"
	fn "marketplacecc/common/functionNames"
	mp "marketplacecc/marketplace"
	dp "marketplacecc/marketplace/data-processor"
)

type MarketPlaceCC struct {
}

/*
Main method to start the chaincode
*/
func main() {
	//127.0.0.1:57051
	//os.Setenv("CORE_PEER_ADDRESS", "127.0.0.1:7052")
	//os.Setenv("CORE_CHAINCODE_ID_NAME", "marketplace:v1")
	//fmt.Println(os.Args)
	err := shim.Start(new(MarketPlaceCC))
	if err != nil {
		log.Panic("error starting MarketPlaceCC chaincode: %s", err.Error())
	}
}

// Init initializes chaincode
func (mpCC *MarketPlaceCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Println("Initializing MarketPlaceCC Chaincode")
	processor := makeCCDataProcessor(stub)
	_, args := stub.GetFunctionAndParameters()
	response := mp.Init(processor, args)
	return response
	return shim.Success(nil)
}

func makeCCDataProcessor(stub shim.ChaincodeStubInterface) *dp.DataProcessor {
	reader := dp.NewCCHLObjectReader(stub)
	writer := dp.NewCCHLObjectWriter(stub, reader)
	return &dp.DataProcessor{
		HlReader: reader,
		HlWriter: writer,
		Stub:     stub,
	}
}

// Invoke initializes chaincode
func (mpCC *MarketPlaceCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	log.Println("invoke is running " + function)
	processor := makeCCDataProcessor(stub)
	switch function {
	case fn.AddAndUpdateCCUser:
		return mp.AddAndUpdateCCUser(processor, args)
		break
	case fn.AddAndUpdateEntity:
		return mp.AddAndUpdateEntity(processor, args)
		break
	case fn.AddAndUpdateUser:
		return mp.AddAndUpdateUser(processor, args)
		break
	case fn.AddAndUpdateLots:
		return mp.AddAndUpdateLots(processor, args)
		break
	case fn.DeleteKey:
		return mp.DeleteKeys(processor, args)
		break
	default:
		break
	}
	log.Println("invoke did not find func: " + function)
	return shim.Error("Received unknown function invocation")
}
