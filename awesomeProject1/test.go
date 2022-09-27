package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type SimpleAsset struct {
}

func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetArgs()
	if len(args) != 2 {
		return shim.Error("args != 2")
	}
	err := stub.PutState(string(args[0]), []byte(args[1]))
	if err != nil {
		return shim.Error("PutState failed")
	} else {
		return shim.Success(nil)
	}
}

func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()
	if fn == "set" {
		result, err := t.Set(stub, args)
		if err != nil {
			return shim.Error("Invoking...failed to set")
		} else {
			return shim.Success([]byte(result))
		}
	} else {
		result, err := t.Get(stub, args)
		if err != nil {
			return shim.Error("Invoking...failed to set")
		} else {
			return shim.Success([]byte(result))
		}
	}
}
func (t *SimpleAsset) Set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("set(arges) args != 2")
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", fmt.Errorf("failed to set %s", args[0])
	} else {
		return args[1], nil
	}
}
func (t *SimpleAsset) Get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("set(arges) args != 2")
	}
	result, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("failed to get %s", args[0])
	}
	return string(result), nil
}
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
