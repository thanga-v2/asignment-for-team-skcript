/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Structure for Diamond

// Red diamond is one of the most expensisve diamond

// To track this diamond from the place it originated and to check the authenticity of the diamond

type Diamond struct {
	Name              string `json:name`
	DateofManufacture string `json:Date_of_Manufacturing`
	Cost              string `json:cost`
	Status            string `json:status`
	Cert              string `json:cert`
	OwnerID           string `json:OwnerID`
	OwnerName         string `json:OwnerName`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments

	//to receive the payload

	//function and arguments will be seperated

	//api stub will be provided by chaincode itself

	//APIStub is a library method, to debuf the pay load

	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "createAsset" {
		return s.createAsset(APIstub, args)
	} else if function == "queryAsset" {
		return s.queryAsset(APIstub, args)
	} else if function == "transferOfOwnership" {
		return s.transferOfOwnership(APIstub, args)
	} else if function == "certifyAsset" {
		return s.certifyAsset(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) createAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	//checking the arguments
	if len(args) != 8 {
		return shim.Error("Expecting 8 arguments to create an asset")
	}

	//go lang object

	var diamond = Diamond{
		Name:              args[1],
		DateofManufacture: args[2],
		Cost:              args[3],
		Status:            args[4],
		Cert:              args[5],
		OwnerID:           args[6],
		OwnerName:         args[7]}

	// to convert the above Golang object into json bytes

	diamondAsBytes, _ := json.Marshal(diamond)

	// here we are writing so -> PutState

	// arg[0] -> identity of the row

	APIstub.PutState(args[0], diamondAsBytes)

	return shim.Success(nil)

}

func (s *SmartContract) queryAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	diamondAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(diamondAsBytes)
}

func (s *SmartContract) certifyAsset(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	// to certify the asset
	// to change the Status to Listed

	if len(args) != 2 {
		shim.Error("Incorrect number of arguments")
	}

	diamondAsBytes, _ := APIstub.GetState(args[0])
	diamond := Diamond{}

	//Unmarshaling to get the bytes and reference
	json.Unmarshal(diamondAsBytes, &diamond)
	diamond.Status = args[1]

	// adding the certificate to the asset
	diamond.Cert = args[2]

	diamondAsBytes, _ = json.Marshal(diamond)
	APIstub.PutState(args[0], diamondAsBytes)

	return shim.Success(nil)

}

// when buyer tries to buy the asset

func (s *SmartContract) transferOfOwnership(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 4 {
		shim.Error("arguments doesnt meet the match")
	}

	diamondAsBytes, _ := APIstub.GetState(args[0])
	diamond := Diamond{}

	//Unmarshaling to get the bytes and reference
	json.Unmarshal(diamondAsBytes, &diamond)

	//checking the amount

	// amount has to be matched and then only the ownership will be changed to new owner

	// converting string to int

	// cost data type has to be converted from string to int to check if the amount is matching or not

	//
	iSeller, _ := strconv.Atoi(diamond.Cost)
	iBuyer, _ := strconv.Atoi(args[1])

	if iSeller == iBuyer {

		//owner ID is the publickey

		//if buyer buys the Diamond, Owner ID of seller has to be transfered to Buyer

		//diamond.OwnerID = args[2]
		diamond.OwnerName = args[2]
	} else {
		return shim.Error("amount for the particular product is not matching")
	}

	diamondAsBytes, _ = json.Marshal(diamond)
	APIstub.PutState(args[0], diamondAsBytes)

	return shim.Success(nil)

}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
