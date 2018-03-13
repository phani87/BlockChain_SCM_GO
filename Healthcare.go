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

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initHealthcare","args":["mer1000001", "mercedes", "c class", "1502688979", "ser1234", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initHealthcare","args":["maz1000001", "mazda", "mazda 6", "1502688979", "ser1235", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initHealthcare","args":["ren1000001", "renault", "megan", "1502688979", "ser1236", "renault", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initHealthcare","args":["ford1000001", "ford", "mustang", "1502688979", "ser1237", "ford", "false", "1502688979"],"chaincodeVer":"v1"}'

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
	//"bytes"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	//"strings"
	//"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// AutoTraceChaincode example simple Chaincode implementation
type HealthCareChainCode struct {
}


// @MODIFY_HERE add recall fields to vehicle JSON object
// type Doctor struct {
// 	ObjectType        string `json:"docType"` 
// 	Doctor_id         string `json:"doctor_id"` 
// 	Doctor_name       string `json:"doctor_name"`
// 	Doctor_speciality string `json:"doctor_speciality"`
// 	Doctor_ph_no       int   `json:"doctor_ph_no"`
	
// }

// type Diagnostic struct {
// 	ObjectType        	string `json:"docType"` 
// 	Icd_code         	string `json:"icd_code"` 
// 	Diagnostic_desc     string `json:"diagnostic_desc"`
// 	Diagnostic_action 	string `json:"diagnostic_action"`	
// }

// type Prescription struct {
// 	ObjectType        			string `json:"docType"` 
// 	Prescription_id         	string `json:"prescription_id"` 
// 	Prescription_medName    	string `json:"prescription_medName"`
// 	Prescription_instruction 	string `json:"prescription_instruction"`	
// }

// type Patient struct {
// 	ObjectType        			string `json:"docType"` 
// 	Patient_id         			string `json:"patient_id"` 
// 	Patient_name       			string `json:"patient_name"`
// 	Patient_dob 				int `json:"patient_dob"`
// 	Patient_gender 				string `json:"patient_gender"`
// 	Patient_weight 				string `json:"patient_weight"`
// 	Patient_bp					string `json:"patient_bp"`
// 	Patient_body_temp 			string `json:"patient_body_temp"`
// 	Patient_ph_no				string `json:"patient_ph_no"`
// 	Patient_zip 				string `json:"patient_zip"`
// }

type Visit struct {
	ObjectType     				string `json:"docType"` 
	Visit_id 					string `json:"visit_id"`
	Doctor_id   				string `json:"doctor_id"` 
	Icd_code    				string `json:"icd_code"`  
	Prescription_id				string `json:"prescription_id"`
	Prescription_instruction 	string `json:"prescription_instruction"`	 
	Patient_id      			string `json:"patient_id"` 
	Visit_date 					int `json:"visit_date"` 
}



// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(HealthCareChainCode))
	if err != nil {
		fmt.Printf("Error starting Parts Trace chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *HealthCareChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *HealthCareChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initHealthCare" { //create new mobile device
		return t.initHealthCare(stub, args)
	}
	// } else if function == "getHistoryForRecord" { //get history of values for a record
	// 	return t.getHistoryForRecord(stub, args)
	// } else if function == "transferMobile" {
	// 	return t.transferMobile(stub, args)
	// }

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMobile - create a new mobile , store into chaincode state
// ============================================================
func (t *HealthCareChainCode) initHealthCare(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// data visit -patient
	//   0       		1      		2     		3			4		   5	   6		7			8
	// "pt123456", "john doe", "1502688979", "male", "171 pounds", "120/75", "98", "335-996-6654", "270358"


		// Visit_id 				
		// Doctor_id   			
		// Icd_code    			
		// Prescription_id			
		// Prescription_instruction
		// Patient_id      		
		// Visit_date 				

	// @MODIFY_HERE extend to expect 6 arguements, up from 6
	if len(args) <= 5  {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init healthcare")
	if len(args[0]) <= 0 {
		return shim.Error("Visit ID cannot be empty")
	}
	if len(args[1]) <= 0 {
		return shim.Error("Doctor id  cannot be empty ")
	}
	if len(args[2]) <= 0 {
		return shim.Error("ICD code cannot be empty ")
	}
	if len(args[3]) <= 0 {
		return shim.Error("Prescription ID cannot be empty")
	}
	if len(args[4]) <= 0 {
		return shim.Error("Prescription Instruction cannot be empty")
	}
	if len(args[5]) <= 0 {
		return shim.Error("Patient id cannot be empty")
	}
	if len(args[6]) <= 0 {
		return shim.Error("Visit Date code cannot be empty ")
	}

	//Set Init parameters

	Visit_id 					:= args[0]
	Doctor_id 					:= args[1]
	Icd_code         			:= args[2]
	Prescription_id         	:= args[3]
	Prescription_instruction	:= args[4]
	Patient_id         			:= args[5]
	Visit_date, err 			:= strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("Visit Date must be a numeric string")
	}


	// ==== Check if Visit already exists ====
	visitAsBytes, err := stub.GetState(Visit_id)
	if err != nil {
		return shim.Error("Failed to get visit : " + err.Error())
	} else if doctorAsBytes != nil {
		return shim.Error("This Visit already exists: " + Visit_id)
	}

	// // ==== Check if patient already exists ====
	// patientAsBytes, err := stub.GetState(patient_id)
	// if err != nil {
	// 	return shim.Error("Failed to get Mobile: " + err.Error())
	// } else if patientAsBytes != nil {
	// 	return shim.Error("This Mobile already exists: " + IMEINumber)
	// }

	// ==== Create Visit object and marshal to JSON ====

	objectType := "visit"
	v := &visit{objectType, Visit_id, Doctor_id, Icd_code, Prescription_id, Prescription_instruction, Patient_id, Visit_date}
	//v := &Visit{Visit_id, Doctor{"Doctor",Doctor_id,Doctor_name,Doctor_speciality,Doctor_ph_no},Diagnostic{"Diagnostic", Icd_code, Diagnostic_desc, Diagnostic_action},Prescription{"Prescription", Prescription_id,Prescription_medName,Prescription_instruction},Patient{"Patient",Patient_id,Patient_name,Patient_dob,Patient_gender,Patient_weight,Patient_bp, Patient_body_temp,Patient_ph_no,Patient_zip} }

	
	visitJSONasBytes, err := json.Marshal(v)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save mobile to state ===
	err = stub.PutState(Visit_id, visitJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	//  ==== Index the mobile based on the owner
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  In our case, the composite key is based on indexName~assember~chassisNumber.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexName := "doctor_id~visit_id"
	ownersIndex := "patient_id~visit_id"
	err = t.createIndex(stub, indexName, []string{v.Doctor_id, v.Visit_id})
	if err != nil {
		return shim.Error(err.Error())
	}
	err = t.createIndex(stub, ownersIndex, []string{v.Patient_id, v.Visit_id})
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== Hospital Visit part saved and indexed. Return success ====
	fmt.Println("- end init Hospital Visit")
	return shim.Success(nil)
}

// ===============================================
// createIndex - create search index for ledger
// ===============================================
func (t *HealthCareChainCode) createIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	fmt.Println("- start create index")
	var err error
	//  ==== Index the object to enable range queries, e.g. return all parts made by supplier b ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of object.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(indexKey, value)

	fmt.Println("- end create index")
	return nil
}


// ===========================================================================================
// cryptoVerify : Verifies signed message against public key
// Public Key of Authority:
// [48 78 48 16 6 7 42 134 72 206 61 2 1 6 5 43 129 4 0 33 3 58 0 4 21 162 242 84 40 78 13 26 160 33 97 191 210 22 152 134 162 66 12 77 221 129 138 60 74 243 198 34 102 209 14 48 16 2 98 96 172 47 170 216 228 169 103 121 153 100 84 111 33 13 106 42 46 227 52 91]
// ===========================================================================================
func cryptoVerify(hash []byte, publicKeyBytes []byte, r *big.Int, s *big.Int) (result bool) {
	fmt.Println("- Verifying ECDSA signature")
	fmt.Println("Message")
	fmt.Println(hash)
	fmt.Println("Public Key")
	fmt.Println(publicKeyBytes)
	fmt.Println("r")
	fmt.Println(r)
	fmt.Println("s")
	fmt.Println(s)

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	switch publicKey := publicKey.(type) {
	case *ecdsa.PublicKey:
		return ecdsa.Verify(publicKey, hash, r, s)
	default:
		return false
	}
}