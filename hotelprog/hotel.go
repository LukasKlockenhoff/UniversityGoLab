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

/* ==================== exercise 2 ==================== */
type hotel struct {
	name     string
	distance int
	exchange bool
	next     *hotel
}


func init_tour() []hotel {
	// create hotels
	var tour []hotel
	fmt.Println("How long should the tour be?")
	var length int
	fmt.Scanln(&length)
	for i := 0; i < length; i++ {
		fmt.Println("Hotel Name: ")
		var name string
		fmt.Scanln(&name)
		fmt.Println("Distance: ")
		var distance int
		fmt.Scanln(&distance)
		fmt.Println("Exchange: ")
		var exchange bool
		fmt.Scanln(&exchange)
		
		tour.append(hotel{name, distance, exchange, nil})
		fmt.Println(tour[i].name, "Confirmed")
	}
	// link hotels
	for i := 0; i < length-1; i++ {
		tour[i].next = &tour[i+1]
	}
	fmt.Println("Tour created")
	// validate tour + add roundtrip
	fmt.Println("Is it a round trip? (y/n)")
	var round_trip string
	fmt.Scanln(&round_trip)
	if round_trip == "y" {
		tour[length-1].next = &tour[0]
	}

	if tour_validate(length, tour) {
		return tour
	}
	fmt.Println("Tour not valid")
	return nil
}

func tour_validate(length int, tour []hotel) bool {
	// check if tour is valid
	if length < 2{
		fmt.Println("Tour too short")
		return false
	}
	for i := 0; i < length-1; i++ {
		if tour[i].next == nil {
			fmt.Println("Hotels not properly linked")
			return false
		}
	}
	fmt.Println("Tour is valid")
	return true
}

/* ==================== exercise 3 ==================== */

func calc_distance(tour []hotel) int {
	// calculate total distance of tour
	var distance int
	for i := 0; i < len(tour); i++ {
		distance += tour[i].distance
	}
	return distance
}

func ride_backwards(tour []hotel) []hotel {
	// ride tour backwards
	var new_tour []hotel
	for i := len(tour)-1; i >= 0; i-- {
		if i == 0 {
			tour[i].next = nil
			new_tour.append(tour[i])
		} else {
			tour[i].next = &tour[i-1]
			new_tour.append(tour[i])
		}
	}
	return new_tour
}

/* ==================== exercise 4 ==================== */

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
	hotels, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hotels)
}

/* ==================== exercise 5 ==================== */

/*
Aufgabe 5: Entwickeln Sie Ihr Hotelprogramm so weiter, dass es wesentliche Funktionen (Laden, Speichern, Wegsuche) in einem Package bereitstellt.

Aufgabe 6: Erweitern Sie Ihr Hotelprogramm um eine Funktion, die eine zufällige Tour berechnet.

Aufgabe 10: Erweitern Sie Ihr Hotelprogramm um eine Funktion, die eine optimale Tour berechnet, die alle Hotels mit Wechselmöglichkeit besucht und dabei die Gesamtdistanz minimiert. Die Funktion soll die Tour in einer Datei speichern.

// TODO: Aufgabe 5: Package hotel mit Funktionen zum Laden, Speichern und Berechnen von Touren
// TODO: Aufgabe 6: Zufällige Tour berechnen -> Funktion in Package hotel, hotel.next = random(hotel)
// TODO: Aufgabe 10: Optimale Tour berechnen -> Funktion in Package hotel, graphen aus hotel erstellen, graphen mit Dijkstra durchlaufen, tour aus graphen erstellen, DFS, BFS, A*, ...
// Hotel als Knoten, Hotel.next als Kante, Distanz als Gewicht -> Hotel.next als Liste von Nachbarn, Distanz als Liste von Gewichten
*/



func main() {
	testfahrt = init_tour()
	calc_distance(testfahrt)

}

