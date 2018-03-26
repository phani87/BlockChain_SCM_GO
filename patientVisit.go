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


// // @MODIFY_HERE add recall fields to vehicle JSON object
// type visit struct {
// 	ObjectType         	string `json:"docType"` 
// 	VisitId         	string `json:"visitId"` 
// 	DoctorId        	string `json:"doctorId"`
// 	DoctorName			string `json:"doctorName"`
// 	DoctorTotBill		string `json:"doctorTotBill"`
// 	DoctorInsPay		string `json:"doctorInsPay"`
// 	DoctorCoPay 		string `json:"doctorCoPay"`
// 	PatientId           string `json:"patientId"`
// 	PatientName 		string `json:"patientName"`
// 	RxId      		 	string `json:"rxId"`
// 	RxDrugs				string `json:"rxDrugs"`
// 	RxInstructions		string `json:"rxInstructions"`
// 	RxTotBill			string `json:"rxTotBill"`
// 	RxInsPay			string `json:"rxInsBill"`
// 	RxCoPay				string `json:"rxCoPay`
// 	InsId				string `json:"insId"`
// 	VisitDate           int    `json:"visitDate"`
// } 

// @MODIFY_HERE add recall fields to vehicle JSON object
type patient struct {
	ObjectType         		string `json:"docType"`
	PatientId           	string 	`json:"patientId"`
	PatientFirstName 		string 	`json:"patientFirstName"`
	PatientLastName			string	`json:"patientLastName"`
	PatientAge				string	`json:"patientAge"`
	PatientBp				string  `json:"patientBp"`
	PatientDob				int		`json:"patientDob"`
} 

// @MODIFY_HERE add recall fields to vehicle JSON object
type prescription struct {
	ObjectType         	string `json:"docType"`
	RxId      		 	string `json:"rxId"`
	RxDrugs				string `json:"rxDrugs"`
	RxInstructions		string `json:"rxInstructions"`
	RxTotBill			string `json:"rxTotBill"`
	RxInsPay			string `json:"rxInsBill"`
	RxCoPay				string `json:"rxCoPay`
	InsId				string `json:"insId"`
	DoctorId        	string `json:"doctorId"`
	DoctorName			string `json:"doctorName"`
	PatientId           string 	`json:"patientId"`
	RxDate				int	   `json:"rxDate"`
}

type visit struct {
	ObjectType         	string `json:"docType"`
	VisitId         	string 	`json:"visitId"` 
	DoctorId        	string 	`json:"doctorId"`
	DoctorName			string	`json:"doctorName"`
	DoctorTotBill		string 	`json:"doctorTotBill"`
	DoctorInsPay		string 	`json:"doctorInsPay"`
	DoctorCoPay 		string 	`json:"doctorCoPay"`
	PatientId           string 	`json:"patientId"`
	VisitNotes			string 	`json:"visitNotes"`
	VisitDate           int    	`json:"visitDate"`
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
	if function == "initPatient" { //create new mobile device
		return t.initPatient(stub, args)
	}else if function == "initVisit" { //create new mobile device
		return t.initVisit(stub, args)
	} else if function == "initPrescription" { //create new mobile device
		return t.initPrescription(stub, args)
	} else if function == "getHistoryForRecord" { //get history of values for a record
		return t.getHistoryForRecord(stub, args)
	} else if function == "getCompostiteKey" { //get history of values for a record
		return t.getCompostiteKey(stub, args)
	}
	
	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initPatient - create a new patient , store into chaincode state
// ============================================================
func (t *PtntVstTraceChaincode) initPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// @MODIFY_HERE extend to expect 5 arguements, up from 6
	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init vehicle")

	if len(args[0]) <= 0 {
		return shim.Error("patient id must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("patient first name must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("patient last name must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("patient age must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("patient BP must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("patient date of birth must be a non-empty string")
	}
	
	patientId			:=args[0]
	patientFirstName	:=strings.ToLower(args[1])
	patientLastName		:=strings.ToLower(args[2])
	patientAge			:=strings.ToLower(args[3])
	patientBp			:=strings.ToLower(args[4])
	patientDob, err 	:= strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("patient date of birth must be a numeric string")
	}			


	// ==== Check if patient already exists ====
	patientAsBytes, err := stub.GetState(patientId)
	if err != nil {
		return shim.Error("Failed to get Visit: " + err.Error())
	} else if patientAsBytes != nil {
		return shim.Error("This patient already exists: " + patientId)
	}

	// @MODIFY_HERE parts recall fields
	// ==== Create vist  and marshal to JSON ====
	objectType := "patient"
	
	patient := &patient{objectType, patientId, patientFirstName, patientLastName, patientAge, patientBp, patientDob}
	patientJSONasBytes, err := json.Marshal(patient)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save mobile to state ===
	err = stub.PutState(patientId, patientJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	//  ==== Index the mobile based on the owner
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on patient~details
	//  This will enable very efficient state range queries based on composite keys matching patient~details~*

	// patientIndex := "patient~details"
	// patientIndexKey, err := stub.CreateCompositeKey(patientIndex, []string{patient.PatientId, patient.PatientFirstName, patient.PatientLastName})
	// if err != nil {
	// 	return shim.Error(err.Error())
	// }
	// //  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	// //  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	// value := []byte{0x00}
	// stub.PutState(patientIndexKey, value)

	// ==== Visit part saved and indexed. Return success ====
	fmt.Println("- end init Visit")
	return shim.Success(nil)
}


// ============================================================
// initPatient - create a new visit , store into chaincode state
// ============================================================
func (t *PtntVstTraceChaincode) initVisit(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// @MODIFY_HERE extend to expect 5 arguements, up from 6
	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
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
		return shim.Error("doctor total bill  must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("patient id must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("visit notes be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("visit date be a non-empty string")
	}
	
	visitId				:=args[0]
	doctorId			:=strings.ToLower(args[1])
	doctorName			:=strings.ToLower(args[2])
	patientId			:=strings.ToLower(args[3])
	visitNotes			:=strings.ToLower(args[4])
	visitDate, err 	:= strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("patient date of birth must be a numeric string")
	}			

	doctorTotBill 		:= ""
	doctorInsPay		:= ""
	doctorCoPay			:= ""

	// ==== Check if patient already exists ====
	visitAsBytes, err := stub.GetState(patientId)
	if err != nil {
		return shim.Error("Failed to get Visit: " + err.Error())
	} else if visitAsBytes != nil {
		return shim.Error("This patient already exists: " + patientId)
	}

	// @MODIFY_HERE parts recall fields
	// ==== Create vist  and marshal to JSON ====
	objectType := "visit"
	
	visit := &visit{objectType, visitId, doctorId, doctorName, doctorTotBill, doctorInsPay, doctorCoPay, patientId, visitNotes, visitDate}
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
	//  In our case, the composite key is based on patient~details
	//  This will enable very efficient state range queries based on composite keys matching patient~details~*

	visitIndex := "patient~visit~doctor"
	visitIndexKey, err := stub.CreateCompositeKey(visitIndex, []string{visit.PatientId, visit.VisitId, visitId.DoctorId})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(visitIndexKey, value)

	// ==== Visit part saved and indexed. Return success ====
	fmt.Println("- end init Visit")
	return shim.Success(nil)
}



// ============================================================
// initPatient - create a new visit , store into chaincode state
// ============================================================
func (t *PtntVstTraceChaincode) initPrescription(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// @MODIFY_HERE extend to expect 5 arguements, up from 6
	if len(args) < 5 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
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
		return shim.Error("doctor total bill  must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("patient id must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("visit notes be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("visit date be a non-empty string")
	}
	
	rxId					:=args[0]
	rxDrugs					:=strings.ToLower(args[1])
	rxInstructions			:=strings.ToLower(args[2])
	insId					:=strings.ToLower(args[3])
	doctorId				:=strings.ToLower(args[4])
	doctorName				:=strings.ToLower(args[5])
	patientId				:=strings.ToLower(args[6])
	rxDate, err 			:= strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("patient date of birth must be a numeric string")
	}			

	rxTotBill 		:= ""
	rxInsBill		:= ""
	rxCoPay			:= ""

	// ==== Check if patient already exists ====
	prescriptionAsBytes, err := stub.GetState(rxId)
	if err != nil {
		return shim.Error("Failed to get Visit: " + err.Error())
	} else if prescriptionAsBytes != nil {
		return shim.Error("This prescription already exists: " + rxId)
	}

	// @MODIFY_HERE parts recall fields
	// ==== Create vist  and marshal to JSON ====
	objectType := "visit"
	
	prescription := &prescription{objectType, rxId, rxDrugs, rxInstructions, rxTotBill, rxInsBill, rxCoPay, insId, doctorId, doctorName, patientId, rxDate}
	prescriptionJSONasBytes, err := json.Marshal(prescription)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save mobile to state ===
	err = stub.PutState(rxId, prescriptionJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	//  ==== Index the mobile based on the owner
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on patient~details
	//  This will enable very efficient state range queries based on composite keys matching patient~details~*

	visitIndex := "patient~rx~doctor"
	visitIndexKey, err := stub.CreateCompositeKey(visitIndex, []string{visit.PatientId, visit.RxId, visitId.DoctorId})
	if err != nil {
		return shim.Error(err.Error())
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of the marble.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(visitIndexKey, value)

	// ==== Visit part saved and indexed. Return success ====
	fmt.Println("- end init Visit")
	return shim.Success(nil)
}


//=================================================================================================================================================================================================

 
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


		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


//============================================================================================
// Searches based on the composite key and returns all the data requested for based on 
// search string and the key like, Search String (patient~rx~doctor, patient~visit~doctor) and 
// Key could be Patient ID or Doctor ID or Visit ID 
// Work around in case, couchDB cannot be configured
//============================================================================================

func (t *PtntVstTraceChaincode) getCompostiteKey(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	keyString := args[0]
	keyToSearch := args[1]
	
	resultsIterator, err := stub.GetStateByPartialCompositeKey(keyString, []string{keyToSearch})
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
		key1 := compositeKeyParts[0]
		key2 := compositeKeyParts[1]
		key3 := compositeKeyParts[2]
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"key1\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(key1))
		buffer.WriteString("\"")

		buffer.WriteString(",\"key2\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(key2))
		buffer.WriteString("\"")

		buffer.WriteString(",\"key3\":")
		buffer.WriteString("\"")
		buffer.WriteString(string(key3))
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



//=================================================================================================================================================================================================

