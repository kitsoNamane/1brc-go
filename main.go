package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Station struct {
	Name  string
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	stations := make(map[string]*Station)
	file, err := os.Open("measurements.txt")
	if err != nil {
		panic("failed to open file")
	}

	csv := csv.NewReader(file)
	csv.Comma = ';'

	for {
		line, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("failed to read line " + err.Error())
			continue
		}

		stationName := line[0]
		temp, _ := strconv.ParseFloat(line[1], 64)
		if _, ok := stations[stationName]; !ok {
			stations[stationName] = &Station{
				Name:  stationName,
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			stations[stationName].Count = stations[stationName].Count + 1
			stations[stationName].Sum = stations[stationName].Sum + temp
			if stations[stationName].Min > temp {
				stations[stationName].Min = temp
			}
			if stations[stationName].Max > temp {
				stations[stationName].Max = temp
			}
		}
	}

	fmt.Print("{")
	for _, station := range stations {
		fmt.Print(fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", station.Name, station.Min, station.Sum/float64(station.Count), station.Max))
	}
	fmt.Println("}")

}
