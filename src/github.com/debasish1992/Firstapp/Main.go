package main

import (
	"fmt"
)

var (
	actorName       = "Choroolet De witte"
	actorProfession = "DJ"
	actorage        = 27
	actorAddress    = "Texas, USA"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	a2 = iota
)

func main() {
	var actorAddress string = "San Jose, CA"
	actorAddress = "Berlin, Germany"
	fmt.Println("Where is Procs", actorAddress)

	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v,  %T\n", n, n)
	fmt.Printf("%v,  %T\n", m, m)

	var a int = 10
	var b int8 = 3
	fmt.Println(a + int(b))

	calculateWithBinaries()

	dealingWithConstants()

	getIotaValues()

}

func calculateWithBinaries() {
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

}

func getIotaValues() {
	fmt.Printf("Valus is %v\n", a)
	fmt.Printf("Value is %v\n", b)
	fmt.Printf("Value is %v\n", c)
	fmt.Printf("Value is %v\n", a2)

}

func bitShifting() {
	a := 8              // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1
}

func runeUsages() {
	r := 'd'
	fmt.Printf("%v, %T\n", r, r)
}

func dealingWithConstants() {
	const myConst int = 99
	fmt.Println("%v, %T\n", myConst, myConst)
}
