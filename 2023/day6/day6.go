package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve())
}

func solve() int {
	raceInfos := getRaceInfos()

	var waysToWin []int

	for index, time := range raceInfos[0] {
		var distances []int
		var v int = 0
		var a int = 1
		for i := 1; i < time; i++ {
			v = a * i
			distance := v * (time - i)
			//fmt.Printf("v: %v | d: %v\n", v, distance)
			if distance > raceInfos[1][index] {
				distances = append(distances, distance)
			}
		}
		//fmt.Println(time)
		//fmt.Println(distances)
		waysToWin = append(waysToWin, len(distances))
	}

	sol := 1
	for _, wTw := range waysToWin {
		sol *= wTw
	}

	return sol
}

func getRaceInfos() [2][]int {
	file, err := os.Open("./inputP2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	var times []int
	for _, time := range strings.Split(strings.Trim(strings.Split(scanner.Text(), ":")[1], " "), " ") {
		if time == "" {
			continue
		}
		timeint, err := strconv.Atoi(time)
		if err != nil {
			log.Fatal(err)
		}
		times = append(times, timeint)
	}

	var distances []int
	scanner.Scan()
	for _, distance := range strings.Split(strings.TrimPrefix(strings.Split(scanner.Text(), ":")[1], " "), " ") {
		if distance == "" {
			continue
		}
		distanceint, err := strconv.Atoi(distance)
		if err != nil {
			log.Fatal(err)
		}
		distances = append(distances, distanceint)
	}

	return [2][]int{
		times,
		distances,
	}
}
