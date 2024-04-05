package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"runtime/trace"
	"strconv"
	"strings"
)

type Station struct {
	Name  string
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	traceFile, err := os.Create("t.out")
	if err != nil {
		panic("failed to create trace file")
	}

	trace.Start(traceFile)
	defer trace.Stop()

	stations := make(map[string]*Station)
	file, err := os.Open("measurements.txt")
	if err != nil {
		panic("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	csv := csv.NewReader(file)
	csv.Comma = ';'

	for scanner.Scan() {
		stationName, tempStr, sepFound := strings.Cut(scanner.Text(), ";")

		if !sepFound {
			continue
		}

		temp, _ := strconv.ParseFloat(tempStr, 64)
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
