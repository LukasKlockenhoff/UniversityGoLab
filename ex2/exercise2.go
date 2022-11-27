package main

import "fmt"

// AUFGABE 1
func one() {
	array := [10]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(array)
}

func two() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 1; i < len(array); i++ {
		fmt.Print(array[i])
	}
	counter := 0
	for counter < len(array) {
		fmt.Print(array[counter])
		counter++
	}

	for i := range array {
		fmt.Print(i)
	}

	i := 0
	for {
		fmt.Print(array)
		i++
		if i == 4 {
			break
		}
	}
}

func three() {
	type user struct {
		username string
		password string
	}

	u1 := user{"Ironman", "123"}
	u2 := user{"CAmerica", "1933"}
	u3 := user{"Hulk", "Strong123"}
	u4 := user{"Thor", "Mjolnir"}
	u5 := user{"Loki", "abc"}

	accounts := [5]user{u1, u2, u3, u4, u5}

	var UName string
	var PW string
	fmt.Print("USERNAME: ")
	fmt.Scanln(&UName)
	fmt.Print("PASSWORD: ")
	fmt.Scanln(&PW)

	for i := 0; i < len(accounts); i++ {
		if accounts[i].username == UName {
			for j := 0; j < 3; j++ {
				if accounts[i].password == PW {
					fmt.Println("LOGIN SUCCESSFUL.")
					return
				} else {
					fmt.Print("PASSWORD: ")
					fmt.Scanln(&PW)
				}
			}
			fmt.Println("LOGIN FAILED")
			return
		}

	}
	fmt.Println("USERNAME NOT FOUND")
	return

}

