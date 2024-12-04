package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Position struct {
	X int
	Y int
}

type Crossword struct {
	Direction string     // left, right, up, down, left_up, left_down, right_up, right_down
	Positions []Position // Array of x and y positions relative to the file (data []string)
}

func main() {
	//fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//mt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

func part1(filename string) int {
	data := readFile(filename)

	var crosswords []Crossword

	re_X := regexp.MustCompile(`X`)

	var index_list_x [][]int

	for i, _ := range data {
		index_list_x = re_X.FindAllStringIndex(data[i], -1)

		for _, x_index := range index_list_x {
			directions := findDirection(data, Position{i, x_index[0]}, 'M')

			for _, direction := range directions {
				x_pos := Position{i, x_index[0]}
				m_pos := findNextCharPosition(data, 'M', x_pos, direction)
				a_pos := findNextCharPosition(data, 'A', m_pos, direction)
				s_pos := findNextCharPosition(data, 'S', a_pos, direction)
				if s_pos.X == -1 || s_pos.Y == -1 || a_pos.X == -1 || a_pos.Y == -1 || m_pos.X == -1 || m_pos.Y == -1 || x_pos.X == -1 || x_pos.Y == -1 {
					continue
				}

				crossword := Crossword{
					Direction: direction,
					Positions: []Position{x_pos, m_pos, a_pos, s_pos},
				}
				if getCrossword(data, crossword) != "XMAS" {
					log.Fatal("Crossword is not XMAS")
				}
				crosswords = append(crosswords, crossword)
			}
		}

	}

	return len(crosswords)
}

func getCrossword(data []string, crossword Crossword) string {
	var word string
	for _, pos := range crossword.Positions {
		word = word + string(data[pos.X][pos.Y])
	}

	return word
}

func findNextCharPosition(data []string, char byte, startPos Position, direction string) Position {
	var nextPosition Position
	switch direction {
	case "left":
		nextPosition = Position{startPos.X, startPos.Y - 1}
	case "right":
		nextPosition = Position{startPos.X, startPos.Y + 1}
	case "up":
		nextPosition = Position{startPos.X - 1, startPos.Y}
	case "down":
		nextPosition = Position{startPos.X + 1, startPos.Y}
	case "left_up":
		nextPosition = Position{startPos.X - 1, startPos.Y - 1}
	case "left_down":
		nextPosition = Position{startPos.X + 1, startPos.Y - 1}
	case "right_up":
		nextPosition = Position{startPos.X - 1, startPos.Y + 1}
	case "right_down":
		nextPosition = Position{startPos.X + 1, startPos.Y + 1}
	}

	if nextPosition.X >= 0 && nextPosition.X < len(data) && nextPosition.Y >= 0 && nextPosition.Y < len(data[0]) {
		if data[nextPosition.X][nextPosition.Y] == char {
			return nextPosition
		}
	}

	return Position{-1, -1}
}

// Searches surrounding fields for the next character
func findDirection(data []string, startPos Position, char byte) []string {
	directions := []string{"left", "right", "up", "down", "left_up", "left_down", "right_up", "right_down"}
	var validDirections []string

	for _, direction := range directions {
		nextPos := findNextCharPosition(data, char, startPos, direction)
		if nextPos.X != -1 && nextPos.Y != -1 {
			validDirections = append(validDirections, direction)
		}
	}

	return validDirections
}

func safeDirection(direction string, pos Position) bool {
	if direction == "left" && pos.Y-1 >= 0 {
		return true
	}
	if direction == "right" && pos.Y+1 < 10 {
		return true
	}
	if direction == "up" && pos.X-1 >= 0 {
		return true
	}
	if direction == "down" && pos.X+1 < 10 {
		return true
	}
	if direction == "left_up" && pos.X-1 >= 0 && pos.Y-1 >= 0 {
		return true
	}
	if direction == "left_down" && pos.X+1 < 10 && pos.Y-1 >= 0 {
		return true
	}
	if direction == "right_up" && pos.X-1 >= 0 && pos.Y+1 < 10 {
		return true
	}
	if direction == "right_down" && pos.X+1 < 10 && pos.Y+1 < 10 {
		return true
	}

	return false
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}
