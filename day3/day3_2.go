package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const SYMBOLS = "*"

func main() {
	fmt.Println("Sum of part numbers:")
	fmt.Println(solve())
}

func solve() int {
	file, err := os.Open("./testinput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var solution int = 0

	var enginePlan [][]string

	var gearPositions [][5]int

	for scanner.Scan() {
		var lineArr []string = strings.Split(scanner.Text(), "")
		enginePlan = append(enginePlan, lineArr)
	}

	lineLength := len(enginePlan[0])

	dotLine := [][]string{strings.Split(strings.Repeat(".", lineLength), "")}

	// add "."s as first and last line
	enginePlan = append(dotLine, enginePlan...)
	enginePlan = append(enginePlan, dotLine...)

	var beginDigitPos int
	var endDigitPos int

	for i := 1; i < len(enginePlan); i++ {
		wasDigit := false
		beginDigitPos, endDigitPos = -1, -1
		for pos, char := range enginePlan[i] {
			if wasDigit == false {
				if _, err := strconv.Atoi(char); err == nil {
					beginDigitPos = pos
					wasDigit = true
				}
				if pos != lineLength-1 {
					continue
				}
			}
			if _, err := strconv.Atoi(char); err == nil {
				endDigitPos = pos
				wasDigit = true
				if pos != lineLength-1 {
					continue
				}
			}
			if beginDigitPos != -1 && endDigitPos == -1 {
				endDigitPos = beginDigitPos
			}
			if beginDigitPos != -1 && endDigitPos != -1 {

				fmt.Println(enginePlan[i][beginDigitPos : endDigitPos+1])
				searchDigitBeginPos := int(math.Max(float64(beginDigitPos-1), 0))
				searchDigitEndPos := int(math.Min(float64(endDigitPos+1), float64(lineLength-1)))
				compareLineBeforeString := strings.Join(enginePlan[i-1][searchDigitBeginPos:searchDigitEndPos+1], "")
				compareLineAfterString := strings.Join(enginePlan[i+1][searchDigitBeginPos:searchDigitEndPos+1], "")
				beforeOrAfter := false
				//fmt.Printf("\nStartPos: %d\nEndPos: %d\nSearchStart: %d\nSearchEnd: %d\n", beginDigitPos, endDigitPos, searchDigitBeginPos, searchDigitEndPos)
				//if strings.Join(enginePlan[i][beginDigitPos:endDigitPos+1], "") == "755" {
				//	fmt.Println(math.Max(float64(beginDigitPos), 0))
				//}
				if beginDigitPos != 0 {
					beforeOrAfter = strings.Contains(SYMBOLS, enginePlan[i][beginDigitPos-1])
				}
				if endDigitPos != len(enginePlan[i])-1 && beforeOrAfter == false {
					beforeOrAfter = strings.Contains(SYMBOLS, enginePlan[i][endDigitPos+1])
				}
				if strings.ContainsAny(compareLineBeforeString, SYMBOLS) ||
					strings.ContainsAny(compareLineAfterString, SYMBOLS) ||
					beforeOrAfter {
					sol, err := strconv.Atoi(strings.Join(enginePlan[i][beginDigitPos:endDigitPos+1], ""))
					if err != nil {
						log.Fatal(err)
					}
					// fmt.Println("Has adjustant *-symbol")
					if strings.ContainsAny(compareLineBeforeString, SYMBOLS) {
						gearCount := strings.Count(compareLineBeforeString, "*")
						fmt.Println("Before Contains Gears: ")
						fmt.Println(gearCount)
						var lastGearXPos int = -1
						for j := 0; j < gearCount; j++ {
							xGearPos := searchDigitBeginPos + strings.Index(compareLineBeforeString[lastGearXPos+1:], "*") + lastGearXPos + 1
							lastGearXPos = strings.Index(compareLineBeforeString[lastGearXPos+1:], "*")
							gearPositions = append(gearPositions, [5]int{xGearPos, i - 1, sol, beginDigitPos, endDigitPos})
							//fmt.Println("Added gear pos:")
							//fmt.Println([2]int{xGearPos, i - 1})
						}
					}
					if strings.ContainsAny(compareLineAfterString, SYMBOLS) {
						gearCount := strings.Count(compareLineAfterString, "*")
						fmt.Println("After Contains Gears: ")
						fmt.Println(gearCount)
						var lastGearXPos int = -1
						for j := 0; j < gearCount; j++ {
							xGearPos := searchDigitBeginPos + strings.Index(compareLineAfterString[lastGearXPos+1:], "*") + lastGearXPos + 1
							lastGearXPos = strings.Index(compareLineAfterString[lastGearXPos+1:], "*")
							gearPositions = append(gearPositions, [5]int{xGearPos, i + 1, sol, beginDigitPos, endDigitPos})
							//fmt.Println("Added gear pos:")
							//fmt.Println([2]int{xGearPos, i + 1})
						}
					}
					if beforeOrAfter {
						before, after := false, false
						if beginDigitPos != 0 {
							before = strings.Contains(SYMBOLS, enginePlan[i][beginDigitPos-1])
							if before {
								gearPositions = append(gearPositions, [5]int{beginDigitPos - 1, i, sol, beginDigitPos, endDigitPos})
							}
						}
						if endDigitPos != len(enginePlan[i])-1 {
							after = strings.Contains(SYMBOLS, enginePlan[i][endDigitPos+1])
							if after {
								gearPositions = append(gearPositions, [5]int{endDigitPos + 1, i, sol, beginDigitPos, endDigitPos})
							}
						}
					}
				}
				wasDigit = false
				beginDigitPos, endDigitPos = -1, -1
			}
		}
	}

	//fmt.Println(gearPositions)
	solution = solveForGears(gearPositions)

	return solution
}

func solveForGears(gearPositions [][5]int) int {
	var possibleGears []string
	var possibleGearsWithoutDigits []string
	var realGears string = ""
	for _, gear := range gearPositions {
		stringArr := [5]string{strconv.Itoa(gear[0]), strconv.Itoa(gear[1]), strconv.Itoa(gear[2]), strconv.Itoa(gear[3]), strconv.Itoa(gear[4])}
		possibleGearsWithoutDigits = append(possibleGearsWithoutDigits, strings.Join(stringArr[:2], ";"))
		possibleGears = append(possibleGears, strings.Join(stringArr[:], ";"))
	}
	fmt.Println(possibleGears)
	fmt.Println(possibleGearsWithoutDigits)
	for index, gear := range possibleGears {
		count := countInSlice(possibleGearsWithoutDigits,
			func(s string) bool { return s == possibleGearsWithoutDigits[index] })
		//count := strings.Count(strings.Join(possibleGearsWithoutDigits, ":"), possibleGearsWithoutDigits[index])
		fmt.Println("Searching for:")
		fmt.Println(possibleGearsWithoutDigits[index])
		fmt.Println("Search in:")
		fmt.Println(strings.Join(possibleGearsWithoutDigits, ":"))
		fmt.Println(count)
		if count == 2 {
			fmt.Println("Found right gear:")
			fmt.Println(gear)
			if !strings.Contains(realGears, gear) {
				realGears = realGears + ":" + gear
			}
		}
	}
	realGears = realGears[1:]

	realGearsArr := strings.Split(realGears, ":")
	sort.Slice(realGearsArr, func(i, j int) bool {
		gearArrI := strings.Split(realGearsArr[i], ";")
		gearArrJ := strings.Split(realGearsArr[j], ";")
		if gearArrI[1] == gearArrJ[1] {
			return gearArrI[0] < gearArrJ[0]
		}
		return gearArrI[1] < gearArrJ[1]
	})
	//fmt.Println(realGearsArr)
	var sol int = 0
	for i := 0; i < len(realGearsArr); i += 2 {
		part1, err := strconv.Atoi(strings.Split(realGearsArr[i], ";")[2])
		if err != nil {
			log.Fatal(err)
		}
		part2, err := strconv.Atoi(strings.Split(realGearsArr[i+1], ";")[2])
		if err != nil {
			log.Fatal(err)
		}
		sol += part1 * part2
	}
	return sol
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
