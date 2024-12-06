package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards    string
	strength int
	bet      int
}

var cardStrengthMap = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

func main() {
	fmt.Println(solve())
}

func solve() int {
	games := getGames()
	var hands []Hand
	for _, game := range games {
		bet, err := strconv.Atoi(game[1])
		if err != nil {
			log.Fatal(err)
		}

		hand := Hand{cards: game[0], bet: bet}
		hand.setStrength()
		//fmt.Printf("Card: %v\nStrength: %v\n\n", hand.cards, hand.strength)
		hands = append(hands, hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].strength == hands[j].strength {
			for index, card := range strings.Split(hands[i].cards, "") {
				if cardStrengthMap[card] == cardStrengthMap[string(hands[j].cards[index])] {
					continue
				}
				return cardStrengthMap[card] < cardStrengthMap[string(hands[j].cards[index])]
			}
			log.Fatal("Found same hands, cannot determine rank")
		}
		return hands[i].strength < hands[j].strength
	})

	rank := 1
	sol := 0
	for _, hand := range hands {
		sol += hand.bet * rank
		rank++
	}

	return sol
}

func getGames() [][2]string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games [][2]string

	for scanner.Scan() {
		game := strings.Split(scanner.Text(), " ")
		games = append(games, [2]string{game[0], game[1]})
	}

	return games
}

func (hand *Hand) setStrength() {
	// five of a kind
	if strings.Count(hand.cards, string(hand.cards[0])) == 5 {
		hand.strength = 7
		return
	}

	// four of a kind
	if strings.Count(hand.cards, string(hand.cards[0])) == 4 ||
		strings.Count(hand.cards, string(hand.cards[1])) == 4 {
		strength := 6
		if strings.Contains(hand.cards, "J") {
			strength = 7
		}
		hand.strength = strength
		return
	}

	// full house
	cardCounts := make(map[string]int)
	for _, card := range hand.cards {
		cardCounts[string(card)]++
	}
	if len(cardCounts) == 2 {
		for _, cardCount := range cardCounts {
			if cardCount == 2 || cardCount == 3 {
				strength := 5
				if strings.Contains(hand.cards, "J") {
					strength = 7
				}
				hand.strength = strength
				return
			}
		}
	}

	// three of a kind
	for _, cardCount := range cardCounts {
		if cardCount == 3 {
			strength := 4
			if strings.Contains(hand.cards, "J") {
				strength = 6
			}
			hand.strength = strength
			return
		}
	}

	// two pair
	pairCnt := 0
	for _, cardCount := range cardCounts {
		if cardCount == 2 {
			pairCnt++
			if pairCnt == 2 {
				strength := 3
				if strings.Contains(hand.cards, "J") {
					strength = 5
					if strings.Count(hand.cards, "J") == 2 {
						strength = 6
					}
				}
				hand.strength = strength
				return
			}
		}
	}

	// one pair
	for _, cardCount := range cardCounts {
		if cardCount == 2 {
			strength := 2
			if strings.Contains(hand.cards, "J") {
				strength = 4
			}
			hand.strength = strength
			return
		}
	}

	// high card
	strength := 1
	if strings.Contains(hand.cards, "J") {
		strength = 2
	}
	hand.strength = strength
}
