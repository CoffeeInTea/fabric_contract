/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	studentcard "StudentCardProject/StudentCard-Contract/student-card"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {

	contract := new(studentcard.StudentCardContract)
	contract.TransactionContextHandler = new(studentcard.TransactionContext)
	contract.Name = "org.studentcontract.studentcard"
	contract.Info.Version = "0.0.1"

	chaincode, err := contractapi.NewChaincode(contract)

	if err != nil {
		panic(fmt.Sprintf("Error creating chaincode. %s", err.Error()))
	}

	chaincode.Info.Title = "StudentCardChaincode"
	chaincode.Info.Version = "0.0.1"

	err = chaincode.Start()

	if err != nil {
		panic(fmt.Sprintf("Error starting chaincode. %s", err.Error()))
	}
}
