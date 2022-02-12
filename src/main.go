package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	Name   string
	Points []string
	Symbol string
}

type Match struct {
	Players   []Player
	CreatedAt time.Time
	Board     [][]string
	Winner    string
}

func (match *Match) RegisterSymbolOnBoard(player int, place []int) {
	playerOnMatch := match.Players[player]
	symbol := playerOnMatch.Symbol

	match.Board[place[0]][place[1]] = symbol
}

func Any(arr []bool) (int, bool) {
	for index, element := range arr {
		if element {
			return index, true
		}
	}
	return -1, false
}

func All(arr []bool) bool {
	for _, element := range arr {
		if !element {
			return false
		}
	}
	return true
}

func StringIn(arr []string, target string) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}

func CompareStringArrayWithMultidimensionalArray(arr [][]string, targets []string) bool {
	checks := []bool{}
	sort.Strings(targets)

	for _, value := range arr {
		sort.Strings(value)
		checks = append(checks, reflect.DeepEqual(targets, value))
	}

	_, ok := Any(checks)
	return ok
}

func CreateMatch() Match {
	var firstPlayerName string
	var secondPlayerName string

	fmt.Print("First player: ")
	fmt.Scan(&firstPlayerName)

	fmt.Print("Second player: ")
	fmt.Scan(&secondPlayerName)

	return Match{
		Players: []Player{
			{
				Name:   firstPlayerName,
				Points: []string{},
				Symbol: "x",
			},
			{
				Name:   secondPlayerName,
				Points: []string{},
				Symbol: "O",
			},
		},
		Board:     [][]string{{"", "", ""}, {"", "", ""}, {"", "", ""}},
		CreatedAt: time.Now(),
		Winner:    "",
	}
}

func GenerateBoard(board [][]string) string {
	stringifiedBoard := ""

	for _, row := range board {
		for _, value := range row {
			stringifiedBoard += fmt.Sprintf("[ %s ]", value)
		}
		stringifiedBoard += "\n"
	}

	return stringifiedBoard
}

func AskForWhereToPutTheSymbol(match Match, player int) string {
	var answer string
	var choices []string

	for rowIndex, row := range match.Board {
		for valueIndex, value := range row {
			if value == "" {
				choices = append(choices, fmt.Sprintf("%d-%d", rowIndex, valueIndex))
			}
		}
	}

	message := fmt.Sprintf("%s, where do you want to put your symbol? Possible places: %s\n>> ", match.Players[player].Name, strings.Join(choices, ", "))
	for answer == "" {
		temporaryAnswer := ""

		fmt.Print(message)
		fmt.Scan(&temporaryAnswer)

		if StringIn(choices, temporaryAnswer) {
			match.Players[player].Points = append(match.Players[player].Points, temporaryAnswer)
			answer = temporaryAnswer
		} else {
			message = "This place is not avaliable! Try another one.\n>> "
		}
	}

	return answer
}

func RegisterSymbolOnBoard(match Match, player int, place []int) {
	playerOnMatch := match.Players[player]
	symbol := playerOnMatch.Symbol

	match.Board[place[0]][place[1]] = symbol
	playerOnMatch.Points = append(playerOnMatch.Points, fmt.Sprintf("%d-%d", place[0], place[1]))

	match.Players[player] = playerOnMatch
}

func main() {
	playerIndex := 0
	possibleWinPatterns := [][]string{
		{
			"0-0",
			"0-1",
			"0-2",
		},
		{
			"1-0",
			"1-1",
			"1-2",
		},
		{
			"2-0",
			"2-1",
			"2-2",
		},
		{
			"0-0",
			"1-0",
			"2-0",
		},
		{
			"0-1",
			"1-1",
			"2-1",
		},
		{
			"0-0",
			"0-1",
			"0-2",
		},
		{
			"0-0",
			"1-1",
			"2-2",
		},
		{
			"0-2",
			"1-1",
			"2-0",
		},
	}
	isGameWon := false

	match := CreateMatch()

	for !isGameWon {
		answer := AskForWhereToPutTheSymbol(match, playerIndex)

		stringifiedIndexes := strings.Split(answer, "-")
		numericIndexes := make([]int, 0)

		for _, index := range stringifiedIndexes {
			result, _ := strconv.Atoi(index)
			numericIndexes = append(numericIndexes, result)
		}
		match.RegisterSymbolOnBoard(playerIndex, numericIndexes)
		board := GenerateBoard(match.Board)
		fmt.Println(board)

		plays := []bool{
			CompareStringArrayWithMultidimensionalArray(possibleWinPatterns, match.Players[0].Points),
			CompareStringArrayWithMultidimensionalArray(possibleWinPatterns, match.Players[1].Points),
		}

		if index, ok := Any(plays); ok {
			match.Winner = match.Players[index].Name
			isGameWon = true
		}

		if playerIndex == 0 {
			playerIndex++
		} else {
			playerIndex--
		}
	}
	fmt.Printf("Player \"%s\" won the game! Thanks for playing! This game took %s.", match.Winner, time.Since(match.CreatedAt))
}
