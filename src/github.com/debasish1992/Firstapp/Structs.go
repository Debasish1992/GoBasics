package main

import "fmt"

type Doctor struct {
	number    int
	actorName string
	companion []string
	episodes  []string
}

func main() {
	aDoctor := Doctor{
		number:    1002,
		actorName: "John Patwee",
		//episodes:  []string{},
		companion: []string{
			"Hayden",
			"Lisa",
			"Simons",
			"Denver",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.companion)

	anonymousStruct()
}

func anonymousStruct() {
	bDoctor := struct{ name string }{name: "Denver Kumar"}
	anotherDoctor := bDoctor
	anotherDoctor.name = "John bakker"
	fmt.Print(bDoctor)
	fmt.Print(anotherDoctor)
}
