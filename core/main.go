package main

import (
	"encoding/json"
	"fmt"

	"github.com/Sedfik/QR_lambgo/core/producer"
)

const (
	Positive bool = true
	Negative      = false
)

type CovidPass struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	DateOfBirth    string `json:"dateOfBirth"`
	CovidStatus    bool   `json:"covidStatus"`
	RequestDate    string `json:"requestDate"`
	GenerationDate string `json:"generationDate"`
}

/*
@return: struct CovidPass
*/
func New(firstName string, lastName string, dateOfBirth string, covidStatus bool) CovidPass {
	return CovidPass{FirstName: firstName, LastName: lastName, DateOfBirth: dateOfBirth}
}

func (covidPass CovidPass) String() string {
	bJson, err := json.Marshal(covidPass)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(bJson)
}

func main() {
	fmt.Println("Hello, world.")
	firstPass := New("Valentin", "Ramos", "26/12/1222", Negative)
	fmt.Println(firstPass)
	producer.Test()
}
