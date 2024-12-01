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

func main() {
	data := sort2D(readFile("input.txt"))
	var results []int

	for _, pair := range data {
		//fmt.Println(pair)
		if pair[0] == pair[1] {
			//fmt.Println(pair[0], "   is equals     ", pair[1] )
			results = append(results, 0)
		}

		if pair[0] > pair[1] {
			//fmt.Println(pair[0], "  is bigger than ", pair[1] )
			results = append(results, pair[0] - pair[1])

		}
		if pair[0] < pair[1] {
			//fmt.Println(pair[0], " is smaller than ", pair[1])
			results = append(results, pair[1] - pair[0])
		}
	}

	var sum int
	for _, result := range results {
		sum = sum + result
	}
	fmt.Println(sum)
}


func sort2D(data [][]int) [][]int {
	rows := len(data)
    cols := len(data[0])

    columns := make([][]int, cols)
    for col := 0; col < cols; col++ {
        columns[col] = make([]int, rows)
        for row := 0; row < rows; row++ {
            columns[col][row] = data[row][col]
        }
    }

    for col := 0; col < cols; col++ {
        sort.Ints(columns[col])
    }

    for col := 0; col < cols; col++ {
        for row := 0; row < rows; row++ {
            data[row][col] = columns[col][row]
        }
    }

	return data
}


func readFile(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Fatal("Unexpected line format: ", line)
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing numbers: ", line)
		}

		data = append(data, []int{num1, num2})
	}

	return data
}



