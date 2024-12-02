package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(part1("example_input.txt"))
	//fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))

}


func part2(filename string) int {
	var similarity_score int

	list1, list2 := readFile(filename)

	//fmt.Println(list1, list2)

	for _, item1 := range list1 {
		var counter int
		for _, item2 := range list2 {
			if item1 == item2 {
				counter++
			}
		}
		similarity_score = similarity_score + item1 * counter
	}

	 

	return similarity_score
}




func part1(filename string) int {
	list1, list2 := readFile(filename)

	slices.Sort(list1)
	slices.Sort(list2)

	var data [][]int

	for i := 0; i < len(list1); i++ {
		data = append(data, []int{list1[i], list2[i]})
	}

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

	return sum
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


func readFile(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int
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

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}



