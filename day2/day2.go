package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func getWinningMoveFromMoveValue(moveValue int) int {
	mod := moveValue % 3

	if mod == 0 {
		return 1
	}

	return mod + 1
}

func getLoosingMoveFromMoveValue(moveValue int) int {
	diff := moveValue - 1

	if diff == 0 {
		return 3
	}

	return diff
}

func getPointsForRoundV1(round string) int {

	oponentMoves := [3]string{
		"A",
		"B",
		"C",
	}

	yourMoves := [3]string{
		"X",
		"Y",
		"Z",
	}

	firstMove := round[:1]
	secondMove := round[2:]

	yourMoveValue := sort.SearchStrings(yourMoves[:], secondMove) + 1
	oponentMoveValue := sort.SearchStrings(oponentMoves[:], firstMove) + 1

	isWin := getWinningMoveFromMoveValue(oponentMoveValue) == yourMoveValue
	isTie := yourMoveValue == oponentMoveValue

	if isWin {
		return yourMoveValue + 6
	}

	if isTie {
		return yourMoveValue + 3
	}

	return yourMoveValue
}

func getPointsForRoundV2(round string) int {
	oponentMoves := [3]string{
		"A",
		"B",
		"C",
	}

	firstMove := round[:1]
	strategy := round[2:]

	oponentMoveValue := sort.SearchStrings(oponentMoves[:], firstMove) + 1

	points := 0

	if strategy == "X" {
		points += getLoosingMoveFromMoveValue(oponentMoveValue)
	}

	if strategy == "Y" {
		points += 3
		points += oponentMoveValue
	}

	if strategy == "Z" {
		points += 6
		points += getWinningMoveFromMoveValue(oponentMoveValue)
	}

	return points

}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	totalPointsPart1 := 0
	totalPointsPart2 := 0
	for fileScanner.Scan() {
		round := fileScanner.Text()
		totalPointsPart1 += getPointsForRoundV1(round)
		totalPointsPart2 += getPointsForRoundV2(round)
	}

	readFile.Close()

	fmt.Println("Part 1:", totalPointsPart1)
	fmt.Println("Part 2:", totalPointsPart2)

}
