package main

import (
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"math"
	"os"
	"strconv"
)

type hotel struct {
	name             string
	distance_to_next int
}

func print(tour []hotel) {
	for _, h := range tour {
		fmt.Println(h.name)
		fmt.Printf("Entfernung zum nächsten Hotel: %dkm\n\n", h.distance_to_next)
	}
}

func read_tour_csv(name string) []hotel {
	//a function that reads a csv file and returns a tour with the hotels and connections
	file, _ := os.Open(name)
	defer file.Close()

	r := csv.NewReader(file)
	records, _ := r.ReadAll()

	var tour []hotel
	var h hotel

	//read the csv file and create a tour

	for _, record := range records {
		h.name = record[0]
		h.distance_to_next, _ = strconv.Atoi(record[1])
		tour = append(tour, h)
	}
	return tour
}

func calc_distance(tour []hotel, start int, end int) int {
	//a function that calculates the distance between two hotels

	// data preparation

	// 1. the user starts counting at 1, but the slice starts at 0

	start = start - 1
	end = end - 1

	// 2. standard case: user enters 1 and 2, 2 and 3, 3 and 4, etc.
	if start < end && end-start == 1 {
		return tour[start].distance_to_next
	}
	// 3. special case: user enters 1 and 5, 2 and 5, 3 and 5, etc.
	if start < end && end-start > 1 {
		var distance int
		for i := start; i <= end; i++ {
			distance += tour[i].distance_to_next
		}
		return distance
	}

	// 4. special case: user enters 5 and 1, 5 and 2, 5 and 3, etc.
	if start > end {
		var distance int
		for i := start; i < len(tour); i++ {
			distance += tour[i].distance_to_next
		}
		for i := 0; i <= end; i++ {
			distance += tour[i].distance_to_next
		}
		return distance
	}
	var distance int

	for i := start; i < end; i++ {
		distance += tour[i].distance_to_next
	}
	return distance
}

// implement a function that uses the bresenham algorithm to draw a line between two points
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm

// Generalized with integer
func Bresenham(img *image.Paletted, x1, y1, x2, y2 int, col uint8) {
	var dx, dy, e, slope int

	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		img.SetColorIndex(x1, y1, col)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			img.SetColorIndex(x1, y1, col)
			x1++
		}
		img.SetColorIndex(x1, y1, col)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			img.SetColorIndex(x1, y1, col)
			y1++
		}
		img.SetColorIndex(x1, y1, col)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				y1--
			}
		}
		img.SetColorIndex(x1, y1, col)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			// BresenhamDxXRYD(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				img.SetColorIndex(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		img.SetColorIndex(x2, y2, col)

	// higher than wide.
	default:
		if y1 < y2 {
			// BresenhamDyXRYD(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.SetColorIndex(x1, y1, col)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				img.SetColorIndex(x1, y1, col)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		img.SetColorIndex(x2, y2, col)
	}
}

// func graphic_represent takes two hotels and represents them as circles. The first hotel green and the second hotel red.
// The distance between the hotels is represented as a line. The function returns a gif. "Hotels.gif"
// The function is used to visualize the distance between the two hotels the user had entered into function calc_distance.
func graphic_represent(h1 hotel, h2 hotel, distance int) {

	// create a new image

	img := image.NewPaletted(image.Rect(0, 0, 1000, 1000), palette.Plan9)

	// draw a white background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// draw a green rectangle
	draw.Draw(img, image.Rect(50, 50, 200, 200), &image.Uniform{color.RGBA{0, 255, 0, 255}}, image.ZP, draw.Src)

	//from the coordinates (200,200) draw a line of length 30 * distance between the hotels using the bresenham algorithm
	// get delta of X and Y:
	dx := int(math.Sqrt((float64(distance*distance) / 2)) * 5)
	dy := dx
	Bresenham(img, 200, 200, 200+dx, 200+dy, 1)

	// draw a red rectangle
	draw.Draw(img, image.Rect(200+dx, 200+dy, (200+dx)+150, (200+dy)+150), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)

	// create a new file
	f, _ := os.Create("Hotels.gif")
	defer f.Close()

	// encode the image as a gif
	gif.Encode(f, img, nil)
}

func main() {
	fmt.Println("Hallo!")
	fmt.Println("Welche Tour möchten Sie laden? Bitte geben Sie den Path zu einer CSV-Datei an.")
	var path string
	fmt.Scanln(&path)

	tour := read_tour_csv(path)

	print(tour)

	fmt.Printf("Sie haben %d Hotels geladen.\n\n\n\n", len(tour))

	for {
		fmt.Println("Was möchten Sie tun?")
		fmt.Println("1 - Distanz zwischen zwei Hotels berechnen")
		fmt.Println("2 - Einen Weg Graphisch darstellen")
		fmt.Println("beliebig - Abbrechen")

		var start int
		var end int
		var input int
		fmt.Scanln(&input)

		if input == 1 {
			fmt.Println("Bitte geben Sie die Nummer des ersten Hotels an.")
			fmt.Scanln(&start)
			fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Println("Bitte geben Sie die Nummer des zweiten Hotels an.")
			fmt.Scanln(&end)
			fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Printf("Die Distanz zwischen Hotel %d und Hotel %d beträgt %dkm.\n", start, end, calc_distance(tour, start, end))
		}
		if input == 2 {
			// make a try to get the user to enter a valid input
			fmt.Println("Welchen Weg möchten Sie darstellen?")
			fmt.Println("Bitte geben Sie die Nummer des ersten Hotels an.")
			fmt.Scanln(&start)
			fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Println("Bitte geben Sie die Nummer des zweiten Hotels an.")
			fmt.Scanln(&end)
			fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			graphic_represent(tour[start-1], tour[end-1], calc_distance(tour, start, end))
			fmt.Println("Graphic 'Hotels.gif' erstellt.")
			fmt.Println("Legende:")
			fmt.Println("Grün: Startpunkt")
			fmt.Println("Rot: Endpunkt")
			fmt.Println("Blau: Weg mit Maßstab 1km = 10px")

		}
		if input != 1 && input != 2 {
			break
		}
	}

}
