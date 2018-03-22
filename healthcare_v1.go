/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

// ====CHAINCODE EXECUTION SAMPLES (BCS REST API) ==================
/*
#TEST transaction / Init ledger

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerB","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerC","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerD","args":["ser1234"],"chaincodeVer":"v1"}'

# TEST transaction / Add Car Part

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1234", "tata", "1502688979", "airbag 2020", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1235", "tata", "1502688979", "airbag 2020", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1236", "tata", "1502688979", "airbag 2020", "toyota", "false", "15026889790"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1237", "tata", "1502688979", "airbag 5000", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1238", "tata", "1502688979", "airbag 5000", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1239", "tata", "1502688979", "airbag 5000", "toyota", "false", "15026889790"],"chaincodeVer":"v1"}'

# TEST transaction / Add Car

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["mer1000001", "mercedes", "c class", "1502688979", "ser1234", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["maz1000001", "mazda", "mazda 6", "1502688979", "ser1235", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["ren1000001", "renault", "megan", "1502688979", "ser1236", "renault", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["ford1000001", "ford", "mustang", "1502688979", "ser1237", "ford", "false", "1502688979"],"chaincodeVer":"v1"}'

# TEST query / Populated database

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"readVehiclePart","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"readVehicle","args":["mer1000001"],"chaincodeVer":"v1"}'

# TEST transaction / Transfer ownership

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"transferVehiclePart","args":["ser1234", "mercedes"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"transferVehicle","args":["mer1000001", "mercedes los angeles"],"chaincodeVer":"v1"}'

# TEST query / Get History

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"getHistoryForRecord","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"getHistoryForRecord","args":["mer1000001"],"chaincodeVer":"v1"}'

# TEST transaction / delete records

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"deleteVehiclePart","args":["ser1235"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"deleteVehicle","args":["maz1000001"],"chaincodeVer":"v1"}'

# TEST transaction / Recall Part

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/{"channel":"channel1","chaincode":"vehiclenet","method":"setPartRecallState","args":["abg1234",true],"chaincodeVer":"v3"}'



CRYPTO
#Sign
go run cryptoHOL.go -s welcome

#Verify
go run cryptoHOL.go -v welcome 23465785510810132448457841429882907809251724155505686786147550387897 10848776947772665661803987914449872333300709981875993855742805426849
*/

// Index for chaincodeid, docType, owner, size (descending order).
// Note that docType, owner and size fields must be prefixed with the "data" wrapper
// chaincodeid must be added for all queries
//
// Definition for use with Fauxton interface
// {"index":{"fields":[{"data.size":"desc"},{"chaincodeid":"desc"},{"data.docType":"desc"},{"data.owner":"desc"}]},"ddoc":"indexSizeSortDoc", "name":"indexSizeSortDesc","type":"json"}
//
// example curl definition for use with command line
// curl -i -X POST -H "Content-Type: application/json" -d "{\"index\":{\"fields\":[{\"data.size\":\"desc\"},{\"chaincodeid\":\"desc\"},{\"data.docType\":\"desc\"},{\"data.owner\":\"desc\"}]},\"ddoc\":\"indexSizeSortDoc\", \"name\":\"indexSizeSortDesc\",\"type\":\"json\"}" http://hostname:port/channelNameGoesHere/_index

// Rich Query with index design doc and index name specified (Only supported if CouchDB is used as state database):
//   peer chaincode query -C channelNameGoesHere -n vehicleParts -c '{"Args":["queryVehiclePart","{\"selector\":{\"docType\":\"vehiclePart\",\"owner\":\"mercedes\"}, \"use_index\":[\"_design/indexOwnerDoc\", \"indexOwner\"]}"]}'

// Rich Query with index design doc specified only (Only supported if CouchDB is used as state database):
//   peer chaincode query -C channelNameGoesHere -n vehicleParts -c '{"Args":["queryVehiclePart","{\"selector\":{\"docType\":{\"$eq\":\"vehiclePart\"},\"owner\":{\"$eq\":\"mercedes\"},\"assemblyDate\":{\"$gt\":1502688979}},\"fields\":[\"docType\",\"owner\",\"assemblyDate\"],\"sort\":[{\"assemblyDate\":\"desc\"}],\"use_index\":\"_design/indexSizeSortDoc\"}"]}'

package main

import (
	"bytes"
	// "crypto/ecdsa"
	// "crypto/x509"
	"encoding/json"
	"fmt"
	//"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// MobileTraceChainCode example simple Chaincode implementation
type PtntVstTraceChaincode struct {
}


// @MODIFY_HERE add recall fields to vehicle JSON object
type visit struct {
	ObjectType         string `json:"docType"` 
	VisitId         string `json:"visitId"` 
	DoctorId       string `json:"doctorId"`
	DoctorName		string `json:"doctorName"`
	PatientId              string `json:"patientId"`
	PatientName 		  string `json:"patientName"`
	RxId      		 string    `json:"rxId"`
	RxDrugs			string `json:"rxDrugs"`
	RxInstructions	string `json:"rxInstructions"`
	VisitDate              int `json:"visitDate"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(PtntVstTraceChaincode))
	if err != nil {
		fmt.Printf("Error starting Parts Trace chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *PtntVstTraceChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *PtntVstTraceChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initVisit" { //create new mobile device
		return t.initVisit(stub, args)
	} else if function == "getHistoryForRecord" { //get history of values for a record
		return t.getHistoryForRecord(stub, args)
	} else if function == "getVisitByRange" { //get history of values for a record
		return t.getVisitByRange(stub, args)
	} else if function == "queryVisitByDoctor" {
		return t.queryVisitByDoctor(stub, args)
	} else if function == "getVisitCompostiteKey" {
		return t.getVisitCompostiteKey(stub, args)
	}
	
	
	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMobile - create a new mobile , store into chaincode state
// ============================================================
func (t *PtntVstTraceChaincode) initVisit(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//{"channel":"docpharmachannel","chaincode":"healthapp_v1","method":"initVisit","args":["vi123456", "dr12345","pt852369", "rx654789", "20180316"],"chaincodeVer":"v5"}


	// @MODIFY_HERE extend to expect 5 arguements, up from 6
	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init vehicle")
	if len(args[0]) <= 0 {
		return shim.Error("visit id must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("doctor id must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("doctor name must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("patient id must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("patient name must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("rx id must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("rx drugs must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return shim.Error("rx instructions must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return shim.Error("visit date must be a non-empty string")
	}

	visitId := args[0]
	doctorId := strings.ToLower(args[1])
	doctorName := strings.ToLower(args[2])
	patientId := strings.ToLower(args[3])
	patientName := strings.ToLower(args[4])
	rxId := strings.ToLower(args[5])
	rxDrugs := strings.ToLower(args[6])
	rxInstructions := strings.ToLower(args[7])
	visitDate, err := strconv.Atoi(args[8])
	if err != nil {
		return shim.Error("visit Date must be a numeric string")
	}
	

	// ==== Check if mobile already exists ====
	visitAsBytes, err := stub.GetState(visitId)
	if err != nil {
		return shim.Error("Failed to get Mobile: " + err.Error())
	} else if visitAsBytes != nil {
		return shim.Error("This visit already exists: " + visitId)
	}

	// @MODIFY_HERE parts recall fields
	// ==== Create vehicle object and marshal to JSON ====
	objectType := "visit"
	
	visit := &visit{objectType, visitId, doctorId, doctorName, patientId, patientName, rxId, rxDrugs, rxInstructions, visitDate}
	visitJSONasBytes, err := json.Marshal(visit)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save mobile to state ===
	err = stub.PutState(visitId, visitJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	//  ==== Index the mobile based on the owner
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~assember~chassisNumber.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	//visitIndexbyDoc := "doctorId~visitId"
	//visitIndexbyPat := "patientId~visitId"
	// docIndexbyVisit := "visitId~doctorId"
	// patIndexbyVisit := "visitId~patientId"
	// err = t.createIndex(stub, visitIndexbyDoc, []string{visit.DoctorId, visit.VisitId})
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }
	// err = t.createIndex(stub, visitIndexbyPat, []string{visit.PatientId, visit.VisitId})
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }
	// err = t.createIndex(stub, docIndexbyVisit, []string{visit.VisitId, visit.DoctorId})
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }
	// err = t.createIndex(stub, patIndexbyVisit, []string{visit.VisitId, visit.PatientId})
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }

	indexName := "doctor~visit~patient"
	DocPatViIndexKey, err := stub.CreateCompositeKey(indexName, []string{visit.DoctorId, visit.VisitId, visit.PatientId})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(DocPatViIndexKey, value)

	// ==== Vehicle part saved and indexed. Return success ====
	fmt.Println("- end init vehicle")
	return shim.Success(nil)
}

// // ===============================================
// // createIndex - create search index for ledger
// // ===============================================
// func (t *PtntVstTraceChaincode) createIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
// 	fmt.Println("- start create index")
// 	var err error
// 	//  ==== Index the object to enable range queries, e.g. return all parts made by supplier b ====
// 	//  An 'index' is a normal key/value entry in state.
// 	//  The key is a composite key, with the elements that you want to range query on listed first.
// 	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
// 	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
// 	if err != nil {
// 		return err
// 	}
// 	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of object.
// 	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
// 	value := []byte{0x00}
// 	stub.PutState(indexKey, value)

// 	fmt.Println("- end create index")
// 	return nil
// }

  
// // ===========================================================================================
// // getHistoryForRecord returns the histotical state transitions for a given key of a record
// // ===========================================================================================
func (t *PtntVstTraceChaincode) getHistoryForRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	recordKey := args[0]

	fmt.Printf("- start getHistoryForRecord: %s\n", recordKey)

	resultsIterator, err := stub.GetHistoryForKey(recordKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the key/value pair
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON vehiclePart)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

//============================================================================================
//Get Visit by range
//============================================================================================

func (t *PtntVstTraceChaincode) getVisitByRange(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

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
		// Add a comma before array members, suppress it for the first array member
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

	fmt.Printf("- getVehiclePartByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (t *PtntVstTraceChaincode) getVisitCompostiteKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	docId := args[0]
	
	resultsIterator, err := stub.GetStateByPartialCompositeKey("doctor~visit~patient", []string{docId})
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
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(queryResponse.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		returnedDocId := compositeKeyParts[0]
		returnedVisitId := compositeKeyParts[1]
		returnedPatientId := compositeKeyParts[1]
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"docId\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(returnedDocId))
		buffer.WriteString("\"")

		buffer.WriteString(",\"visitId\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(returnedVisitId))
		buffer.WriteString("\"")

		buffer.WriteString(",\"patientId\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(returnedPatientId))
		buffer.WriteString("\"")

		buffer.WriteString(",\"objectType\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(objectType))
		buffer.WriteString("\"")

		// buffer.WriteString(", \"Record\":")
		// // Record is a JSON object, so we write as-is
		// buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getVehiclePartByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ===========================================================================================
// Query Based results
// ===========================================================================================

func (t *PtntVstTraceChaincode) queryVisitByDoctor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "bob"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	doctorId := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"visit\",\"doctorId\":\"%s\"}}", doctorId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===========================================================================================
// Query Based results - return metod with resultset
// ===========================================================================================

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
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

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}
