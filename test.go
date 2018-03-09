package main

import (
	"encoding/json"
	"fmt"
)

type Doctor struct {
	ObjectType        string `json:"docType"` 
	Doctor_id         string `json:"doctor_id"` 
	Doctor_name       string `json:"doctor_name"`
	Doctor_speciality string `json:"doctor_speciality"`
	Doctor_ph_no       int   `json:"doctor_ph_no"`
	
}

type Diagnostic struct {
	ObjectType        	string `json:"docType"` 
	Icd_code         	string `json:"icd_code"` 
	Diagnostic_desc     string `json:"diagnostic_desc"`
	Diagnostic_action 	string `json:"diagnostic_action"`	
}

type Visit struct {
	Doctor 		
	Diagnostic 	
}

func main() {

	v := &Visit{Doctor{"Doctor","12345ssdf","John Doe","Cardiologist",3245688452},Diagnostic{"Diagnostic", "asdfa213", "sadfasdfas", "asdfasdf"}}	
	
	js, _ := json.Marshal(v)
	fmt.Printf("%s", js)
}