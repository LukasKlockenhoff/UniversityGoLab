package main

import (
	"fmt"
	"math"
	"strconv"
)

// Task1
const PI = 3.1415

var pi_float float32 = PI

func func_task1() {
	pi_inferr := PI
	fmt.Println(PI)
	fmt.Println(pi_float)
	fmt.Println(pi_inferr)
}

func typecasting(value int16) {
	fmt.Println(int16(value))
	fmt.Println(int32(value))
	var value32 int32 = int32(value)
	fmt.Println(value32)
	fmt.Println(int16(value32))
}

func string_handling() {

	var hallo string = "Hallo Welt"
	var jib string = "ÄÖÜäöü\\"

	fmt.Printf("%s, %d\n", hallo, len(hallo))
	fmt.Printf("%s, %d\n", jib, len(jib))

	//slices ?: init als array
	slice1 := make([]string, len(hallo))
	slice2 := make([]string, len(jib))

	slice3 := []rune(hallo)
	slice4 := []rune(jib)

	fmt.Printf("%s, %d\n", slice1, len(slice1))
	fmt.Printf("%s, %d\n", slice2, len(slice2))

	fmt.Printf("%c, %d\n", slice3, len(slice3))
	fmt.Printf("%c, %d\n", slice4, len(slice4))

	//slicing
	fmt.Println(hallo[3:6])

}

func transformation() {
	zahl := 3000
	//str := "Hi"

	fmt.Printf("Int to string %s\n", strconv.Itoa(zahl))

	zahls := strconv.Itoa(zahl)
	zahld, err := strconv.Atoi("30n")

	fmt.Printf("ERROR: string to int: %d%d", zahld, err) //interesting output

	fmt.Println(zahls, err)

	fmt.Println(strconv.Atoi(zahls))

	fmt.Println(strconv.ParseInt("40", 10, 64)) //Es gibt nich parseBool, float

	fmt.Println(strconv.FormatUint(123, 10)) //Base -> Zahlensystem

}

func arith() {
	if 3 == math.Sqrt(3.0)*math.Sqrt(3.0) {
		fmt.Println("success")
	} else {
		fmt.Println(math.Sqrt(3.0) * math.Sqrt(3.0))
	}

	if 1.0 == 1.0/49*49 {
		fmt.Println("success")
	} else {
		fmt.Println(1.0 == 1.0/49*49)
	}
}

/*
func main() {
	fmt.Println("Hello World!")
	//func_task1()
	//typecasting(30000)
	//string_handling()
	//transformation()
	arith()
}
*/
