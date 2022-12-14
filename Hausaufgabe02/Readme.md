# Hausaufgabe 02

## Aufgabe 1

Entwickeln Sie ein Programm, dass eine Liste von Hotels bestehend aus Namen und Entfernung zum nächsten Hotel aus einer CSV-Datei einliest.

### Meine Implementierung:

Die Funktion read_tour_csv() liest die CSV ein und verteilt die Werte an Structs vom Typ Hotel. Diese werden in einem Array gespeichert und Anschließend zurückgegeben.

### Ausgabefunktion

Die Funktion print(tour []hotel) gibt die Tour formattiert aus.

### Code:

```go
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

func print(tour []hotel) {
	for _, h := range tour {
		fmt.Println(h.name)
		fmt.Printf("Entfernung zum nächsten Hotel: %dkm\n\n", h.distance_to_next)
	}
}
```

## Aufgabe 2
Das Programm fragt den Benutzer nach 2 Hotel-Namen und berechnet die Distanz zwischen den Hotels. Achtung: Der Weg kann auch von einem Hotel mit einer höheren Nummer zu einem Hotel mit kleinerer Nummer erfolgen.

### Meine Implementierung:

Die Funktion calc_distance() berechnet die Distanz zwischen zwei Hotels. Dabei wird die Distanz zwischen den Hotels und die Distanz zwischen den Hotels und dem Startpunkt addiert. Wenn das Hotel mit der größeren Nummer vor dem mit der kleineren Nummer genannt wird, so verwende ich 2 for Schleifen um die Distanz zu berechnen.

### Code:

Zu Beginn werden die Hotels von 1-Idexing in 0-Indexing umgewandelt.
Danach werden verschiedene Fälle abgefragt:
Ob das Hotel mit der kleineren Nummer vor dem Hotel mit der größeren Nummer liegt.
Ob das Hotel mit der kleineren Nummer direkt vor dem Hotel mit der größeren Nummer liegt.
Ob das Hotel mit der kleineren Nummer vor dem Hotel mit der größeren Nummer liegt und die Distanz zwischen den Hotels größer als 1 ist.
...

```go
func calc_distance(tour []hotel, start int, end int) int {
	start = start - 1
	end = end - 1

	if start < end && end-start == 1 {
		return tour[start].distance_to_next
	}

	if start < end && end-start > 1 {
		var distance int
		for i := start; i <= end; i++ {
			distance += tour[i].distance_to_next
		}
		return distance
	}
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
```

## Aufgabe 3
Das Programm berechnet zu dem in Aufgabe 2 berechneten Weg ein Gif-Bild und schreibt dieses in die Datei: „Hotels.gif“
Das Bild enthält die betroffenen Hotels und den Weg zwischen ihnen.
Hinweis: Linien, keine Schrift.

### Meine Implementierung:

Hierzu habe ich ihre Bresenham Funktion verwendet und im Kontext so eingesetz, dass das erste Hotel als grünes Quadrat und das zweite Hotel als rotes Quadrat dargestellt wird. Die Linie zwischen den Hotels wird als blaue Linie dargestellt. Die Länge der Linie ist nicht willkürlich gewählt, sondern entspricht der Distanz zwischen den Hotels.

### Code:

```go
func graphic_represent(h1 hotel, h2 hotel, distance int) {

    //neues bild erstellen
	img := image.NewPaletted(image.Rect(0, 0, 1000, 1000), palette.Plan9)

	// weisser hintergrund
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	// erstes hotel als grünes quadrat
	draw.Draw(img, image.Rect(50, 50, 200, 200), &image.Uniform{color.RGBA{0, 255, 0, 255}}, image.ZP, draw.Src)

	//linie mit bresenham
	// get delta of X and Y:
	dx := int(math.Sqrt((float64(distance*distance) / 2)) * 5)
	dy := dx
	Bresenham(img, 200, 200, 200+dx, 200+dy, 1)

	// zweites hotel als rotes quadrat
	draw.Draw(img, image.Rect(200+dx, 200+dy, (200+dx)+150, (200+dy)+150), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)

	// neue Datei
	f, _ := os.Create("Hotels.gif")
	defer f.Close()

	// als gif speichern
	gif.Encode(f, img, nil)
}
```