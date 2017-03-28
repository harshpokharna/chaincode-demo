package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ChaincodeDemo struct {
}

func main() {
	err := shim.Start(new(ChaincodeDemo))
	if err != nil {
		fmt.Printf("Error starting ChaincodeDemo: %s", err)
	}
}

func (t *ChaincodeDemo) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("ChaincodeDemo.Init()")
	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	err := stub.PutState("message", []byte("Hello, World!"))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *ChaincodeDemo) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("ChaincodeDemo.Invoke() -> " + function)

	if function == "put" {
		if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting 1")
		}

		err := stub.PutState("message", []byte(args[0]))
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation")
}

func (t *ChaincodeDemo) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("ChaincodeDemo.Query() -> " + function)

	if function == "get" {
		if len(args) != 0 {
			return nil, errors.New("Incorrect number of arguments. Expecting 0")
		}
		return stub.GetState("message")
	}

	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query")
}
