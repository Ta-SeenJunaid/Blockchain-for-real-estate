// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Flat structure, with 4 properties.  
Structure tags are used by encoding/json library
*/
type Flat struct {
	Condition string `json:"condition"`
	Ranking string `json:"ranking"`
	Location  string `json:"location"`
	Holder  string `json:"holder"`
}

/*
 * The Init method *
 called when the Smart Contract "flat-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "flat-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryFlat" {
		return s.queryFlat(stub, args)
	} else if function == "initLedger" {
		return s.initLedger(stub)
	} else if function == "recordFlat" {
		return s.recordFlat(stub, args)
	} else if function == "queryAllFlat" {
		return s.queryAllFlat(stub)
	} else if function == "changeFlatHolder" {
		return s.changeFlatHolder(stub, args)
	}else if function == "changeFlatCondition" {
		return s.changeFlatCondition(stub, args)
	}else if function == "changeFlatRanking" {
		return s.changeFlatRanking(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryFlat method *
Used to view the records of one particular flat
It takes one argument -- the key for the flat in question
 */
func (s *SmartContract) queryFlat(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	flatAsBytes, _ := stub.GetState(args[0])
	if flatAsBytes == nil {
		return shim.Error("Could not locate flat")
	}
	return shim.Success(flatAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 flat)to our network
 */
func (s *SmartContract) initLedger(stub shim.ChaincodeStubInterface) sc.Response {
	flat := []Flat{
		Flat{Condition: "923F", Location: "67.0006, -70.5476", Ranking: "1504054225", Holder: "Marjan"},
		Flat{Condition: "M83T", Location: "91.2395, -49.4594", Ranking: "1504057825", Holder: "Som"},
		Flat{Condition: "T012", Location: "58.0148, 59.01391", Ranking: "1493517025", Holder: "Helal"},
		Flat{Condition: "P490", Location: "-45.0945, 0.7949", Ranking: "1496105425", Holder: "Jaman"},
		Flat{Condition: "S439", Location: "-107.6043, 19.5003", Ranking: "1493512301", Holder: "Rafa"},
		Flat{Condition: "J205", Location: "-155.2304, -15.8723", Ranking: "1494117101", Holder: "Shen"},
		Flat{Condition: "S22L", Location: "103.8842, 22.1277", Ranking: "1496104301", Holder: "Leila"},
		Flat{Condition: "EI89", Location: "-132.3207, -34.0983", Ranking: "1485066691", Holder: "Yuan"},
		Flat{Condition: "129R", Location: "153.0054, 12.6429", Ranking: "1485153091", Holder: "Carlo"},
		Flat{Condition: "49W4", Location: "51.9435, 8.2735", Ranking: "1487745091", Holder: "Fatima"},
	}

	i := 0
	for i < len(flat) {
		fmt.Println("i is ", i)
		flatAsBytes, _ := json.Marshal(flat[i])
		stub.PutState(strconv.Itoa(i+1), flatAsBytes)
		fmt.Println("Added", flat[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordflat method *
Fisherman like Sarah would use to record each of her flat catches. 
This method takes in five arguments (attributes to be saved in the ledger). 
 */
func (s *SmartContract) recordFlat(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var flat = Flat{ Condition: args[1], Location: args[2], Ranking: args[3], Holder: args[4] }

	flatAsBytes, _ := json.Marshal(flat)
	err := stub.PutState(args[0], flatAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record flat: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllFlat method *
allows for assessing all the records added to the ledger(all flat)
This method does not take any arguments. Returns JSON string containing results. 
 */
func (s *SmartContract) queryAllFlat(stub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllFlat:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The changeFlatHolder method *
The data in the world state can be updated with who has possession. 
This function takes in 2 arguments, flat id and new holder name. 
 */
func (s *SmartContract) changeFlatHolder(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	flatAsBytes, _ := stub.GetState(args[0])
	if flatAsBytes == nil {
		return shim.Error("Could not find flat")
	}
	flat := Flat{}

	json.Unmarshal(flatAsBytes, &flat)
	// Normally check that the specified argument is a valid holder of flat
	// we are skipping this check for this example
	flat.Holder = args[1]

	flatAsBytes, _ = json.Marshal(flat)
	err := stub.PutState(args[0], flatAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change flat holder: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) changeFlatCondition(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	flatAsBytes, _ := stub.GetState(args[0])
	if flatAsBytes == nil {
		return shim.Error("Could not find flat")
	}
	flat := Flat{}

	json.Unmarshal(flatAsBytes, &flat)
	// Normally check that the specified argument is a valid holder of flat
	// we are skipping this check for this example
	flat.Condition = args[1]

	flatAsBytes, _ = json.Marshal(flat)
	err := stub.PutState(args[0], flatAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change flat condition: %s", args[0]))
	}

	return shim.Success(nil)
}

func (s *SmartContract) changeFlatRanking(stub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	flatAsBytes, _ := stub.GetState(args[0])
	if flatAsBytes == nil {
		return shim.Error("Could not find flat")
	}
	flat := Flat{}

	json.Unmarshal(flatAsBytes, &flat)
	// Normally check that the specified argument is a valid holder of flat
	// we are skipping this check for this example
	flat.Ranking = args[1]

	flatAsBytes, _ = json.Marshal(flat)
	err := stub.PutState(args[0], flatAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change flat Ranking: %s", args[0]))
	}

	return shim.Success(nil)
}



/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}