package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func create_file(name string) {
	//takes a file name and creates a file with that name
	file, err := os.Create(name)
	if err == nil {
		fmt.Println("Datei erfolgreich erstellt.")
		file.WriteString("Hallo")
		file.Close()
	}
}

func write_int(name string, n int) {
	//takes a file name and an integer and writes the integer n times to the file
	file, _ := os.Create(name)
	for i := 0; i < n; i++ {
		file.WriteString(strconv.Itoa(i))
	}
	file.Close()
}

func write_bool(name string) {
	//takes a file name and writes the boolean values true and false to the file
	file, _ := os.Create(name)
	for i := 0; i < 5; i++ {
		file.WriteString(strconv.FormatBool(true))
		file.WriteString(strconv.FormatBool(false))
	}
	file.Close()
}

func write_int_float(name string, integer int64, float float64) {
	// takes a file name, an integer and a float and writes them to the file
	file, _ := os.Create(name)
	for i := 0; i < 10; i++ {
		file.WriteString(strconv.FormatInt(integer, 10))
		file.WriteString(strconv.FormatFloat(float, 64, 10, 64))
	}
	for i := -256; i < 256; i++ {
		file.WriteString(strconv.Itoa(i))
	}
	file.Close()
}

func write_strings(name string, word string) {
	//takes a file name and a string and writes the string 10 times to the file
	file, _ := os.Create(name)
	for i := 0; i < 10; i++ {
		file.WriteString(word)
	}
}

func read_file(name string) {
	//takes a file name and reads the file
	file, _ := os.Open(name)
	fmt.Println(file)
	file.Close()
}

func lottery(name string, date string) {
	//takes a file name and appends randomly generated numbers and a date to the file
	file, _ := os.Create(name)
	//generate random numbers
	for i := 0; i < 7; i++ {
		random := rand.Intn(42)
		file.WriteString(strconv.Itoa(random) + ",")
	}
	//generate date
	file.WriteString(date)
	file.Close()
}

func find_lottery(name string, date string) {
	//takes a file name and searches for the date in the file
	file, _ := os.Open(name)

	defer file.Close()

	//Split the file into lines
	scanner := bufio.NewScanner(file)

	//Loop over the lines
	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), date) {
			fmt.Println(scanner.Text())
		}
		line++
	}

}

func save_lottonumbers_csv(name string, numbers [][]int) {
	//takes a file name and an array of integers and writes the numbers to the file
	file, _ := os.Create(name)
	defer file.Close()

	w := csv.NewWriter(file)
	//write the numbers to the file
	var data [][]string

	for _, value := range numbers {
		row := []string{strconv.Itoa(value[0]), strconv.Itoa(value[1]), strconv.Itoa(value[2]), strconv.Itoa(value[3]), strconv.Itoa(value[4]), strconv.Itoa(value[5]), strconv.Itoa(value[6])}
		data = append(data, row)

	}
	w.WriteAll(data)

}

// Hotels
type hotel struct {
	name     string
	location int
	exchange bool
	next     *hotel
}

func four() []hotel {

	hotel4 := hotel{"TopNotch", 1500, true, nil}
	hotel3 := hotel{"Medium", 400, true, &hotel4}
	hotel2 := hotel{"Luxury", 100, false, &hotel3}
	hotel1 := hotel{"Budget", 0, true, &hotel2}

	tour := []hotel{hotel1, hotel2, hotel3, hotel4}
	tour_validate(4, tour)
	return tour
}

func tour_validate(length int, tour []hotel) {
	if tour[0].exchange && tour[length-1].exchange {
		fmt.Println("TOUR VALID")
		fmt.Println(tour[0].next.name)
		return
	}
	fmt.Println("TOUR INVALID")

}

func tour_csv(name string, tour []hotel) {
	//takes a file name and an array of integers and writes the numbers to the file
	file, _ := os.Create(name)
	defer file.Close()

	w := csv.NewWriter(file)
	//write the numbers to the file
	var data [][]string

	for _, value := range tour {
		row := []string{value.name, strconv.Itoa(value.location), strconv.FormatBool(value.exchange)}
		data = append(data, row)

	}
	w.WriteAll(data)
}

func read_tour_csv(name string) {
	//takes a file name and reads the file
	file, _ := os.Open(name)
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)
}
func main() {
	/*
		create_file("file1.txt")
		write_int("file2.txt", 10)
		write_bool("file3.txt")
		write_int_float("file4.txt", 10, 10.5)
		write_strings("file5.txt", "Hallo")
		read_file("file2.txt")
	*/

	/*
		lottery("lottery.txt", "2020-12-24")
		find_lottery("lottery.txt", "2020-12-24")
		save_lottonumbers_csv("lotto.csv", [][]int{{1, 2, 3, 4, 5, 6, 7}, {1, 2, 3, 4, 5, 6, 7}, {1, 2, 3, 4, 5, 6, 7}})
	*/

	tour := four()
	tour_csv("tour.csv", tour)
	read_tour_csv("tour.csv")
}
