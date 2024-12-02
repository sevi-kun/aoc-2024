package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))

}

func part1(filename string) int {
	data := readFile(filename)
	var safe_report_count int

	for _, report := range data {
		if verifyOrder(report, 0, "unknown") {
			safe_report_count++
		}
	}

	return safe_report_count
}


//The levels are either all increasing or all decreasing.
//Any two adjacent levels differ by at least one and at most three.
func verifyOrder(report []int, index int, orderType string) bool {
	if index+1 == len(report) {
		return true
	}
	if report[index] < report[index+1] && (report[index+1] - report[index]) <= 3 && (orderType == "unknown" || orderType == "up") {
		if verifyOrder(report, index+1, "up") {
			return true
		}
	}
	if report[index] > report[index+1] && (report[index] - report[index+1]) <= 3 && (orderType == "unknown" || orderType == "down") {
		if verifyOrder(report, index+1, "down") {
			return true
		}
	}

	return false
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

		var numLine []int
		for _, numStr := range parts {
			num, err := strconv.Atoi(numStr)
			numLine = append(numLine, num)

			if err != nil {
				log.Fatal("Error parsing numbers: ", line)
			}
		}

		data = append(data, numLine)
	}

	return data
}
