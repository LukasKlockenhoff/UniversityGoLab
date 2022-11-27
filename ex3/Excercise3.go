package main

import "fmt"

//Aufgabe0
func product_table(start int, end int) {
	for i := start; i <= end; i++ {
		for j := start; j <= end; j++ {
			fmt.Print("[", j*i, "]")
		}
		fmt.Print("\n")
	}
}

//Aufgabe1
func print_greater(z1 float64, z2 float64) {
	if z1 > z2 {
		fmt.Println(z1)
	} else if z2 > z1 {
		fmt.Println(z2)
	} else {
		fmt.Println("Beide Zahlen sind gleich gross.")
	}
}

//Aufgabe2
func return_greater(z1 float64, z2 float64) float64 {
	if z1 > z2 {
		return z1
	} else if z2 > z1 {
		return z2
	} else {
		return 0
	}
}

//Aufgabe3
func switch_numbers(n1 float64, n2 float64) (float64, float64) {
	buffer := n1
	n1 = n2
	n2 = buffer
	return n1, n2

}

//Aufgabe4
func iterative_facu(num int) int {
	fac := 1
	for i := 2; i <= num; i++ {
		fac *= i
	}
	return fac
}

//Aufgabe5
func recursive_facu(num int) int {
	if num == 1 {
		return 1
	} else {
		return recursive_facu(num-1) * num
	}
}

//Aufgabe6
///////////////
//Aufgabe7
func rec_fib(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 1
	} else {
		return rec_fib(n-1) + rec_fib(n-2)
	}
}
func it_fib(n int) int {
	fib := 0
	if n == 1 {
		fib = 1
	} else if n == 2 {
		fib = 2
	} else {
		//start at 3
		prev := 1
		prev_prev := 1
		for i := 3; i <= n; i++ {
			fib = prev + prev_prev
			prev_prev = prev
			prev = fib
		}
	}
	return fib
}

//Aufgabe8
type n_und_quad struct {
	n    int
	quad int
}

func array() [100000]n_und_quad {
	one := n_und_quad{1, 1}
	array := [100000]n_und_quad{one}
	for i := 1; i < 100000; i++ {
		array[i] = n_und_quad{i, i * i}
	}
	return array
}

//Aufgabe9
func proc_array(array [100000]n_und_quad) {
	array[0] = n_und_quad{1, 1}
}

//Aufgabe10
//Hotels

func main() {
	//product_table(1, 12)
	//print_greater(1235, 235235.2342)
	//return_greater(1225, 124.1241)
	//fmt.Println(switch_numbers(1, 2))
	//fmt.Println(recursive_facu(4))
	//fmt.Println(rec_fib(7))
	//fmt.Println(it_fib(12))
	//array()
	/*
		for i := 0; i < 10000; i++ {
			proc_array(array())
		}
		fmt.Println("DONE")
	*/
}
