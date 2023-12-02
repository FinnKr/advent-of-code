package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(solve())
}

func solve() int {
	file, err := os.Open("./day1input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var solution int = 0

	for scanner.Scan() {
		var line string = scanner.Text()

		firstDigit, err := findFirstDigitInString(line)
		if err != nil {
			log.Fatal(err)
		}

		runes := []rune(line)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		var reversedLine string = string(runes)

		lastDigit, err := findLastDigitInString(reversedLine)
		if err != nil {
			log.Fatal(err)
		}

		lineSol, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatal(err)
		}
		solution += lineSol
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return solution
}

func findLastDigitInString(s string) (string, error) {
	digitStrings := [9]string{"eno", "owt", "eerht", "ruof", "evif", "xis", "neves", "thgie", "enin"}
	return findFirstDigitInStringSetArray(s, digitStrings[:])
}

func findFirstDigitInString(s string) (string, error) {
	digitStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return findFirstDigitInStringSetArray(s, digitStrings[:])
}

func findFirstDigitInStringSetArray(s string, digitStrings []string) (string, error) {
	var firstWrittenDigitPos int = -1
	var firstWrittenDigitString string

	for index, element := range digitStrings {
		i := strings.Index(s, element)
		if i != -1 && (i < firstWrittenDigitPos || firstWrittenDigitPos == -1) {
			firstWrittenDigitPos = i
			firstWrittenDigitString = strconv.Itoa(index + 1)
		}
	}

	if firstWrittenDigitPos == 0 {
		return firstWrittenDigitString, nil
	}

	var firstDigitPos int = -1
	var firstDigitString string

	for pos, char := range s {
		if unicode.IsDigit(char) {
			firstDigitPos, firstDigitString = pos, string(char)
			break
		}
	}

	if firstDigitPos == -1 && firstWrittenDigitPos == -1 {
		return "", errors.New("No digit found")
	}

	if firstDigitPos == -1 {
		return firstWrittenDigitString, nil
	}

	if firstWrittenDigitPos == -1 {
		return firstDigitString, nil
	}

	if firstDigitPos > firstWrittenDigitPos {
		return firstWrittenDigitString, nil
	}

	return firstDigitString, nil
}
