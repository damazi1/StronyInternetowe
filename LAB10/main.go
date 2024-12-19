package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type car struct {
	mpg          float64 // spalanie (miles per galon)
	cylinders    int     // liczba cylindrów
	displacement float64 // pojemność
	horsepower   float64 // moc
	weight       float64 // masa
	acceleration float64 // przyśpieszenie
	year         int     // rocznik
	origin       int     // pochodzenie
	name         string  // nazwa
}

// funkcja porównująca obiekty przekazane przez parametry
func compare(first *car, second *car) float64 {
	if first == second {
		return 0.0
	}
	similarity := 1.0
	similarity *= 1.0 - math.Abs(first.mpg-second.mpg)/40
	similarity *= 1.0 - math.Abs(first.horsepower-second.horsepower)/300
	similarity *= 1.0 - math.Abs(first.weight-second.weight)/5000
	similarity *= 1.0 - math.Abs(first.acceleration-second.acceleration)/30
	return similarity
}

// funkcja jak metoda - porównuje obiekt bieżący z podanym jako parametr
func (this *car) compare(other *car) float64 {
	if this == other {
		return 0.0
	}
	similarity := 1.0
	similarity *= 1.0 - math.Abs(this.mpg-other.mpg)/40
	similarity *= 1.0 - math.Abs(this.horsepower-other.horsepower)/300
	similarity *= 1.0 - math.Abs(this.weight-other.weight)/5000
	similarity *= 1.0 - math.Abs(this.acceleration-other.acceleration)/30
	return similarity
}

func (this *car) findSim(other []*car) (float64, int) {
	max := 0.0
	indeks := 0
	for i := 0; i < len(other); i++ {
		if this != other[i] {
			if this.compare(other[i]) > max {
				max = this.compare(other[i])
				indeks = i
			}
		}
	}
	return max, indeks
}

func loadCars() []*car {
	var cars []*car
	file, err := os.Open("LAB10/cars.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "\t")
		c := car{}
		c.mpg, _ = strconv.ParseFloat(line[0], 64)
		c.cylinders, _ = strconv.Atoi(line[1])
		c.displacement, _ = strconv.ParseFloat(line[2], 64)
		c.horsepower, _ = strconv.ParseFloat(line[3], 64)
		c.weight, _ = strconv.ParseFloat(line[4], 64)
		c.acceleration, _ = strconv.ParseFloat(line[5], 64)
		c.year, _ = strconv.Atoi(line[6])
		c.origin, _ = strconv.Atoi(line[7])
		c.name = line[8]
		cars = append(cars, &c)
	}
	return cars
}

func main() {
	car0 := car{18, 8, 307, 130, 3504, 12, 70, 1, "chevrolet malibu"}
	car1 := car{13, 8, 351, 158, 4363, 13, 73, 1, "ford ltd"}
	car2 := car{29, 4, 98, 83, 2219, 16.5, 74, 2, "audi fox"}
	car3 := car{20, 6, 232, 100, 2914, 16, 75, 1, "amc gremlin"}
	car4 := car{33, 4, 91, 53, 1795, 17.4, 76, 3, "honda civic"}
	car5 := car{23.2, 4, 156, 105, 2745, 16.7, 78, 1, "plymouth sapporo"}
	cars := []*car{&car0, &car1, &car2, &car3, &car4, &car5}
	fmt.Println(cars)

	cars[4].name = "test"
	fmt.Println(cars[4])
	fmt.Println(car4)

	max := 0.0
	dla := 0
	for i := 0; i < len(cars); i++ {
		if i != 2 {
			if cars[2].compare(cars[i]) > max {
				max = cars[2].compare(cars[i])
				dla = i
			}
		}
	}

	fmt.Println("Najpodobniejszy do pojazdu 2 jest pojazd o numerze", dla, "podobienstwo na poziomie", max)

	cars1 := loadCars()

	fmt.Println(cars1)

	for i := 0; i < len(cars1); i++ {
		if i != 3 {
			if cars1[3].compare(cars1[i]) > max {
				max = cars1[3].compare(cars1[i])
				dla = i
			}
		}
	}
	fmt.Println("Najpodobniejszy do pojazdu 3 jest pojazd o numerze", dla, "podobienstwo na poziomie", max)

	max, indeks := cars1[3].findSim(cars1)

	fmt.Println("Najpodobniejszy do pojazdu 0 jest pojazd o numerze", indeks, "podobienstwo na poziomie", max)

	sam1 := 0
	sam2 := 0
	praw := 0.0

	for i := 0; i < len(cars1); i++ {
		for j := i + 1; j < len(cars1); j++ {
			if cars1[j].compare(cars1[i]) > praw {
				praw = cars1[j].compare(cars1[i])
				sam1 = j
				sam2 = i
			}
		}
	}

	fmt.Println("Najlepsze prawdopodobienstwo bylo dla aut ", sam1, "i", sam2, "o podobienstwie na poziomie", praw)

}
