package main

import (
	"fmt"
)

type resultStructure struct {
	a int
	b int
}

// Function responsible for getting result of a value as 40 from the given array and
// store the indexes into a slice
func main() {
	//var s []int
	var s1 []int
	givenArray := []int{32, 40, 90, 24, 8, 50, 38, 2, 35, 5}
	result := 40
	fmt.Println("Here is the start")
	//dice := resultStructure{100, 200}
	fmt.Println("The Given Array %v", givenArray)
	for i := 0; i < len(givenArray); i++ {
		fetchedInt := givenArray[i]

		fetchedResult := result - fetchedInt

		resultIndex := Contains(givenArray, fetchedResult)

		if fetchedInt == result {
			fmt.Println("This the result on index ", i, fetchedInt)
		} else if resultIndex != -1 {
			findIndex(givenArray, fetchedResult)
			s1 = append(s1, i, resultIndex)
		}
	}
	printSlice(s1)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func contains(a []int, x int) int {
	var resultIndex int = -1
	for i, n := range a {
		//fmt.Println(n)
		if x == n {
			resultIndex = i
		}
	}
	return resultIndex
}

func findIndex(a []int, x int) int {
	var foundIndex = -1
	for i, n := range a {
		if x == n {
			foundIndex = i
		}
	}
	return foundIndex
}
