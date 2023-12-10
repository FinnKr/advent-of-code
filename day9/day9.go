package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE_PATH string = "./input.txt"

func main() {
	fmt.Println(solve())
}

func solve() int {
	datasets := getDataSets()
	sol := 0
	for _, dataset := range datasets {
		//fmt.Println(dataset)
		forecast := getForecast(dataset)
		//fmt.Println(forecast)
		sol += forecast
	}
	return sol
}

func getForecast(dataset []int) int {
	var firstDiffs []int
	currentDataset := dataset
outer:
	for true {
		var differences []int
		for index, data := range currentDataset {
			if index+1 == len(currentDataset) {
				break
			}
			differences = append(differences, currentDataset[index+1]-data)
		}
		//fmt.Println(differences)

		for _, diff := range differences {
			if diff != 0 {
				firstDiffs = append(firstDiffs, differences[0])
				//fmt.Println(firstDiff)
				currentDataset = differences
				continue outer
			}
		}
		break
	}

	substract := firstDiffs[0]
	if len(firstDiffs) > 1 {
		substract = firstDiffs[len(firstDiffs)-2] - firstDiffs[len(firstDiffs)-1]
	}

	for i := len(firstDiffs) - 3; i >= 0; i-- {
		substract = firstDiffs[i] - substract
	}

	return dataset[0] - substract
}

func getDataSets() [][]int {
	file, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var records [][]int

	for scanner.Scan() {
		datasetStr := strings.Split(scanner.Text(), " ")
		var dataset []int
		for _, dsstr := range datasetStr {
			data, err := strconv.Atoi(dsstr)
			if err != nil {
				log.Fatal(err)
			}
			dataset = append(dataset, data)
		}
		records = append(records, dataset)
	}
	return records
}
