// package main

// import "fmt"

// // type Doctor struct {
// // 	number    int
// // 	actorName string
// // 	companion []string
// // 	episodes  []string
// // }

// func main() {
// 	// aDoctor := Doctor{
// 	// 	number:    1002,
// 	// 	actorName: "John Patwee",
// 	// 	//episodes:  []string{},
// 	// 	companion: []string{
// 	// 		"Hayden",
// 	// 		"Lisa",
// 	// 		"Simons",
// 	// 		"Denver",
// 	// 	},
// 	// }

// 	// fmt.Println(aDoctor)
// 	// fmt.Println(aDoctor.actorName)
// 	// fmt.Println(aDoctor.number)
// 	// fmt.Println(aDoctor.companion)

// 	// anonymousStruct()
// 	//solveLooping()
// 	//deferImple()

// 	panicImpl()
// }

// func anonymousStruct() {
// 	// bDoctor := struct{ name string }{name: "Denver Kumar"}
// 	// anotherDoctor := bDoctor
// 	// anotherDoctor.name = "John bakker"
// 	// fmt.Print(bDoctor)
// 	// fmt.Print(anotherDoctor)
// }

// func solveLooping() {
// 	i := 1
// 	for i < 10 {
// 		if i == 9 {
// 			break
// 		}

// 		for j := 0; j < 10; j++ {
// 			fmt.Println(i * j)
// 		}
// 	}
// }

// func deferImple() {

// 	// defer keyword statement execute at kast of the function
// 	defer fmt.Println("Start")
// 	defer fmt.Println("Middle")
// 	defer fmt.Println("End")

// 	a := "start"
// 	// This line will print start not end
// 	defer fmt.Println(a)
// 	a = "end"
// }

// func panicImpl() {
// 	defer fmt.Println("Start")
// 	defer fmt.Println("Middle")
// 	defer fmt.Println("End")
// }
