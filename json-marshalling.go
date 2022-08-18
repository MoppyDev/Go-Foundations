package main

import (
	"encoding/json"
	"log"
)

type Service struct {
	ServiceName  string `json:"service_name"`
	Running      bool   `json:"running"`
	StartupType  string `json:"startup_type"`
	Dependencies string `json:"dependencies"`
}

func main() {
	services := `
	[
		{
			"service_name": "Service1",
			"running": false,
			"startup_type": "Manual",
			"dependencies": "MyService2"
		},
		{
			"service_name": "Service2",
			"running": true,
			"startup_type": "Automatic",
			"dependencies": ""
		}
	]`

	// taking json and assigning it to a struct
	var unmarshalled []Service
	err := json.Unmarshal([]byte(services), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling", err)
	}

	log.Println(unmarshalled)

	// taking struct objects and assigning back into json
	var newServices []Service

	var myNewService Service

	myNewService.ServiceName = "RandomServiceName"
	myNewService.Running = true
	myNewService.StartupType = "Automatic"

	var myNewService2 Service

	myNewService2.ServiceName = "MyCoolService"
	myNewService2.Running = true
	myNewService2.StartupType = "Automatic"

	newServices = append(newServices, myNewService)
	newServices = append(newServices, myNewService2)

	// marshalIndent formats the json with nicer formatting - uncomment line 61 and comment out 60 to compare
	json, err := json.MarshalIndent(newServices, "", "    ")
	//json, err := json.Marshal(newServices)
	if err != nil {
		log.Println("Error Marshalling", err)
	}
	log.Println(string(json))

}
