package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	inputStr string
	id       int
	draws    [][3]int // red, green, blue
}

func (gm *Game) parseString() {
	id, err := strconv.Atoi(gm.inputStr[5:strings.Index(gm.inputStr, ":")])
	if err != nil {
		log.Fatal(err)
	}
	gm.id = id
	strs := strings.Split(gm.inputStr[strings.Index(gm.inputStr, ":")+2:len(gm.inputStr)], "; ")
	gm.draws = make([][3]int, len(strs))

	for i := 0; i < len(strs); i++ {
		gm.draws[i] = [3]int{0, 0, 0}
		draws := strings.Split(strs[i], ", ")
		for j := 0; j < len(draws); j++ {
			var rgbIndex int = 0 // default red
			if strings.Contains(draws[j], "green") {
				rgbIndex = 1
			}
			if strings.Contains(draws[j], "blue") {
				rgbIndex = 2
			}
			amount, err := strconv.Atoi(strings.Split(draws[j], " ")[0])
			if err != nil {
				log.Fatal(err)
			}
			gm.draws[i][rgbIndex] = amount
		}
	}
}

func main() {
	fmt.Println("Sum of minimum powers:")
	fmt.Println(solve())
}

func solve() int {
	file, err := os.Open("./day2input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var solution int = 0

	for scanner.Scan() {
		var line string = scanner.Text()

		game := Game{
			inputStr: line,
		}

		game.parseString()
		minRed, minGreen, minBlue := 0, 0, 0
		for i := 0; i < len(game.draws); i++ {
			if game.draws[i][0] > minRed {
				minRed = game.draws[i][0]
			}
			if game.draws[i][1] > minGreen {
				minGreen = game.draws[i][1]
			}
			if game.draws[i][2] > minBlue {
				minBlue = game.draws[i][2]
			}
		}
		solution += minRed * minGreen * minBlue
	}

	return solution
}
