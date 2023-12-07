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
	fmt.Println(getRaceInfos())
	return -1
}

func getRaceInfos() [2][]int {
	file, err := os.Open("./testinput.txt")
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
