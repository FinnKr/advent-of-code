package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(solve())
}

func solve() int {
	solution := 0
	cards := getCards()
	initCardsCount := len(cards)
	cardCounts := make([]int, initCardsCount)
	for index := range cardCounts {
		cardCounts[index] = 1
	}
	for index, card := range cards {
		cardHits := 0
		for _, drawnNmbr := range card[0] {
			if slices.Contains(card[1], drawnNmbr) {
				cardHits++
			}
		}
		for j := 0; j < cardCounts[index]; j++ {
			for i := 0; i < cardHits; i++ {
				countIndex := index + i + 1
				if countIndex == initCardsCount {
					break
				}
				cardCounts[countIndex] += 1
			}
		}
	}
	for _, count := range cardCounts {
		solution += count
	}
	return solution
}

func getCards() [][2][]int {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cardArray [][2][]int

	for scanner.Scan() {
		allNmbrs := strings.Split(strings.Split(scanner.Text(), ":")[1], " | ")
		drawnNmbrs := strings.Split(allNmbrs[0], " ")
		winNmbrs := strings.Split(allNmbrs[1], " ")

		cardArray = append(cardArray, [2][]int{cleanNmbrsAndConvertToInts(drawnNmbrs), cleanNmbrsAndConvertToInts(winNmbrs)})
	}
	return cardArray
}

func cleanNmbrsAndConvertToInts(nmbrs []string) []int {
	var cleanedNmbrs []int = []int{}
	for _, nmbr := range nmbrs {
		trimmedNmbr := strings.Trim(nmbr, " ")
		if trimmedNmbr != "" {
			nmbrAsInt, err := strconv.Atoi(nmbr)
			if err != nil {
				log.Fatal(err)
			}
			cleanedNmbrs = append(cleanedNmbrs, nmbrAsInt)
		}
	}
	return cleanedNmbrs
}

func countInSlice[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}
