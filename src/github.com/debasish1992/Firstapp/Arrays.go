package main

import "fmt"

func main() {
	// fmt.Println("Learning Arrays")

	// // Declaration of an array
	// grades := [3]int{97, 53, 35}

	// var students [10]string

	// gradesNew := [...]int{97, 53, 35}
	// fmt.Printf("Grades: %v", grades)
	// fmt.Printf("Grades: %v", gradesNew)

	// students[0] = "Lisa"
	// students[1] = "tensa"

	// fmt.Printf("\n Students %v\n", students)
	// fmt.Printf("\n Students Array Length %v\n", len(students))

	//var verifyIdentity bool = playingWithArray()

	// if verifyIdentity {
	// 	fmt.Println("Identity Matched")
	// } else {
	// 	fmt.Println("Identity did not Matched")
	// }

	//playingWithSlices()
	//manipultingSlices()
	creatingSlicesUsingMake()
}

func playingWithArray() bool {
	var isSelectionMatched bool
	var deliveryModelArray [10]string
	valueToMatch := "Sayali"
	deliveryModelArray[0] = "Rama"
	deliveryModelArray[1] = "Shiva"
	deliveryModelArray[2] = "Harini"
	deliveryModelArray[3] = "Sayali"

	for i := 0; i < len(deliveryModelArray); i++ {
		fmt.Println("The Value in the Index is ", deliveryModelArray[i])

		if deliveryModelArray[i] == "Harini" {
			fmt.Println("No Need To Continue")
			break
		}
	}

	// if deliveryModelArray[0] == "Rama" {
	// 	fmt.Printf("This is rama")
	// 	isSelectionMatched = true
	// }

	switch valueToMatch {
	case deliveryModelArray[0]:
		fmt.Println("Found in the 0 index")
	case deliveryModelArray[1]:
		fmt.Println("Found in the 1 index")
	case deliveryModelArray[2]:
		fmt.Println("Found in the 2 index")
	case deliveryModelArray[3]:
		isSelectionMatched = true
		fmt.Println("Found in the 3 index")

	}

	playArray := [...]int{10, 20, 30}
	fmt.Printf("The Value in play array is %v\n", playArray)
	playingArray := &playArray
	playingArray[1] = 99
	fmt.Printf("The Value in Playing Array is %v\n", playingArray)

	return isSelectionMatched

}

func playingWithSlices() {
	mySlice := []int{10, 40, 90, 40, 30}
	// fmt.Printf("The Value Of the Slice is %v\n ", mySlice)
	// fmt.Printf("The length Of the Slice is %v\n ", len(mySlice))
	// fmt.Printf("The capacity Of the Slice is %v\n ", cap(mySlice))
	mySliceTwo := &mySlice
	mySlice[2] = 1000
	fmt.Print(mySliceTwo)
	fmt.Print(mySlice)
}

func manipultingSlices() {
	myDemoSlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// This will Print all the element of the array
	myDemoSlice1 := myDemoSlice[:]
	// This will print the element till the 3rd index
	myDemoSlice2 := myDemoSlice[:3]
	// This will print the elements from 6th element
	myDemoSlice3 := myDemoSlice[6:]
	// This will print the elements from 3rd to 6th
	myDemoSlice4 := myDemoSlice[3:6]
	fmt.Print(myDemoSlice)
	fmt.Print(myDemoSlice1)
	fmt.Print(myDemoSlice2)
	fmt.Print(myDemoSlice3)
	fmt.Print(myDemoSlice4)

	fmt.Print(myDemoSlice4)
}

func creatingSlicesUsingMake() {
	// This will create a slice with length 10 and capacity 100
	a := make([]int, 10, 100)
	fmt.Println(a)
	fmt.Printf("The Array size is %v\n", len(a))
	fmt.Printf("The capacity size is %v\n", cap(a))

	mergingTwoSlices()
}

func mergingTwoSlices() {
	// Blank Slice
	a := []int{}
	// Appending the Element to the Slice
	a = append(a, 2)
	// Printing the elements
	fmt.Print(a)

	// Appening another slice with the prvious slice with the ... operators
	a = append(a, []int{10, 50, 38, 80}...)

	fmt.Print(a)
}
