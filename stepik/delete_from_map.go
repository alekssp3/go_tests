package main

import "fmt"

var groupCity = map[int][]string{
	10:   []string{"Деревня", "Село"}, // города с населением 10-99 тыс. человек
	100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
	1000: []string{"Миллионик"}, // города с населением 1000 тыс. человек и более
 }
 
var cityPopulation = map[string]int{
	"Село" : 100,
	"Миллионик" : 500,
	"Город": 200,
 }

func main() {
	for c, _ := range cityPopulation {
		var needDelete = true
		for _, gc := range groupCity[100] {
			if gc == c {
				needDelete = false
				break
			}
		}
		if needDelete {
			delete(cityPopulation, c)
		}
	}
	fmt.Print(cityPopulation)
}

// for _, city := range append(groupCity[10], groupCity[1000]...) {
//     delete(cityPopulation, city)
// }