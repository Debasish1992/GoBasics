package main

import "fmt"

func main() {
	mapsManipulations()
}

func mapsManipulations() {
	fmt.Printf("This is the starting of the Maps \n")

	// This is the way how we will create a map using make()
	populationMap := make(map[string]int)

	// The Map declaration with String and Int
	populationMap = map[string]int{
		"California": 1209832,
		"Texas":      3218621,
		"Michigan":   2187212,
		"LA":         43243243,
	}

	fmt.Println(populationMap)

	// Adding element to Map
	populationMap["Georgia"] = 1090807060

	// fmt.Printf("The Population in LA is %v\n", populationMap["LA"])
	// fmt.Printf("The Population in texas is %v\n", populationMap["Texas"])
	// fmt.Printf("The Population in Georgia is %v\n", populationMap["Georgia"])

	fmt.Println(populationMap)

	// Deleting elements from the Map
	delete(populationMap, "Georgia")

	fmt.Println("Map After Delete ", populationMap)

	// This ok can be used to check for the element availability in the Map
	// It returns 0, false
	pop, ok := populationMap["Mic"]

	// It returns 2187212, true
	pop2, ok := populationMap["Michigan"]

	fmt.Println(pop, ok)

	fmt.Println(pop2, ok)

	fmt.Println("The Array Size is ", len(populationMap))

	sp := populationMap

	delete(sp, "LA")

	fmt.Println(populationMap)
}
