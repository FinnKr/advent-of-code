package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const INPUT_FILE_PATH string = "./testinput.txt"

func main() {
	fmt.Println(solve())
}

func solve() int {
	return -1
}

func getData() {
	file, err := os.Open(INPUT_FILE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
