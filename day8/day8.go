package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const INPUT_FILE_PATH string = "./input.txt"

type Networkmap map[string][2]string

func main() {
	fmt.Println(solve())
}

func solve() int {
	instructions, networkmap, nodeLabels := readMap()

	//fmt.Println(instructions)
	//fmt.Println(networkmap)
	var stepsArr []int

	for _, nodeLabel := range nodeLabels {
		steps := 0
		for i := 0; i < len(instructions); i++ {
			steps++
			switch instructions[i] {
			case "L":
				nodeLabel = networkmap[nodeLabel][0]
			case "R":
				nodeLabel = networkmap[nodeLabel][1]
			}
			if i == len(instructions)-1 {
				i = -1
			}
			if nodeLabel[len(nodeLabel)-1:] != "Z" {
				continue
			}
			stepsArr = append(stepsArr, steps)
			break
		}
	}

	return LCM(stepsArr[0], stepsArr[1], stepsArr[2:]...)
}

func readMap() ([]string, Networkmap, []string) {
	file, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := strings.Split(scanner.Text(), "")
	scanner.Scan()

	networkmap := make(Networkmap)
	var startingNodes []string

	for scanner.Scan() {
		lineArr := strings.Split(scanner.Text(), " = ")
		valueArr := strings.Split(lineArr[1][1:len(lineArr[1])-1], ", ")
		networkmap[lineArr[0]] = [2]string{valueArr[0], valueArr[1]}
		if lineArr[0][len(lineArr[0])-1:] == "A" {
			startingNodes = append(startingNodes, lineArr[0])
		}
	}

	return instructions, networkmap, startingNodes
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
