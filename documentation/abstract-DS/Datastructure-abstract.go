package main

import "fmt"

type Diamond struct {

	Name string `json:name`
	DateofManufacture string `json:Date_of_Manufacturing`
	Cost int `json:cost`
	Status string `json:status`
	certUrl string `json:certificate_URL`
	OwnerID string `json:OwnerID`
	OwnerName string `json:OwnerName`
}


func main () {
	result_check := Diamond{
		Name:              "The centenary Diamond",
		DateofManufacture: "04-24-1988",
		Cost:              350000,
		Status:            "PENDING",
		certUrl:           "http://localhost:2888/archive/DD/",
		OwnerID:           "OW115",
		OwnerName:         "Winston",
	}

	fmt.Println(result_check)
}
