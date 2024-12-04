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
	//fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

// find X-MAS crosswords (MAS in the shape of an X)
func part2(filename string) int {
	data := readFile(filename)
	var result int

	re_A := regexp.MustCompile(`A`)

	for i, line := range data {

		index_list_a := re_A.FindAllStringIndex(line, -1)

		for _, a_index := range index_list_a {
			a_pos := Position{i, a_index[0]}
			if findXmas(data, a_pos) {
				result++
			}
		}

	}

	return result
}

// Returns true if X-MAS crossword is found at the given a_pos, false otherwise
func findXmas(data []string, a_pos Position) bool {
	var validityScore int // needs exactly 2 matches
	dirs_m := []string{"left_up", "left_down", "right_up", "right_down"}
	dirs_s := []string{"right_down", "right_up", "left_down", "left_up"}

	for i, dir_m := range dirs_m {
		if safeDirection(data, dir_m, a_pos) && safeDirection(data, dirs_s[i], a_pos) {
			if findXmasChar(data, a_pos, dir_m, 'M') && findXmasChar(data, a_pos, dirs_s[i], 'S') {
				validityScore++
			}
		}
	}

	if validityScore == 2 {
		return true
	}
	return false
}

func findXmasChar(data []string, a_pos Position, direction string, char byte) bool {
	nextPos := findNextCharPosition(data, char, a_pos, direction)
	if nextPos.X != -1 && nextPos.Y != -1 {
		return true
	}

	return false
}

// find XMAS crosswords
func part1(filename string) int {
	data := readFile(filename)

	var crosswords []Crossword

	re_X := regexp.MustCompile(`X`)

	var index_list_x [][]int

	for i, _ := range data {
		index_list_x = re_X.FindAllStringIndex(data[i], -1)

		for _, x_index := range index_list_x {
			dirs := []string{"left", "right", "up", "down", "left_up", "left_down", "right_up", "right_down"}
			directions := findDirection(data, dirs, Position{i, x_index[0]}, 'M')

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
func findDirection(data []string, directions []string, startPos Position, char byte) []string {
	var validDirections []string

	for _, direction := range directions {
		nextPos := findNextCharPosition(data, char, startPos, direction)
		if nextPos.X != -1 && nextPos.Y != -1 {
			validDirections = append(validDirections, direction)
		}
	}

	return validDirections
}

func safeDirection(data []string, direction string, pos Position) bool {
	min_x := 0
	min_y := 0
	max_x := len(data)
	max_y := len(data[pos.X])

	if direction == "left" && pos.Y-1 >= min_y {
		return true
	}
	if direction == "right" && pos.Y+1 < max_y {
		return true
	}
	if direction == "up" && pos.X-1 >= min_x {
		return true
	}
	if direction == "down" && pos.X+1 < max_x {
		return true
	}
	if direction == "left_up" && pos.X-1 >= min_x && pos.Y-1 >= min_y {
		return true
	}
	if direction == "left_down" && pos.X+1 < max_x && pos.Y-1 >= min_y {
		return true
	}
	if direction == "right_up" && pos.X-1 >= min_x && pos.Y+1 < max_y {
		return true
	}
	if direction == "right_down" && pos.X+1 < max_x && pos.Y+1 < max_y {
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
