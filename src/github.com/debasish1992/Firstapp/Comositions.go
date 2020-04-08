package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	doctorNumber int
	doctorName   string
}

type Patient struct {
	patientName    string
	patientDeasies string
	patentAddress  string
	Doctor
}

func main() {
	patient := Patient{patientName: "Southee",
		patientDeasies: "Covid19",
		Doctor:         Doctor{doctorName: "CBK Mohanty", doctorNumber: 73}}

	fmt.Println("The Whole Compositon is ", patient)

	goStructTags()

}

func goStructTags() {

	type Employee struct {
		// Tag for the Data
		name string `required max:"100"`
		id   int
	}

	t := reflect.TypeOf(Employee{})
	field, _ := t.FieldByName("name")
	fmt.Println(field.Tag)

}
