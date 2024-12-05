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

func part1(filename string) (result int) {
	rules, updates := readFile(filename)

	fmt.Println(rules)
	fmt.Println(updates)
	
	return 
}

func readFile(filename string) (rules, updates [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "|") {
			before, errb := strconv.Atoi(strings.Split(scanner.Text(), "|")[0])
			after, erra := strconv.Atoi(strings.Split(scanner.Text(), "|")[1])

			if errb != nil || erra != nil {
				log.Fatal(errb, erra)
			}
			rules = append(rules, []int{before, after})
		}
		if strings.Contains(scanner.Text(), ",") {
			numlist := strings.Split(scanner.Text(), ",")
			var nums []int
			for _, num := range numlist {
				n, err := strconv.Atoi(num)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, n)
			}
			updates = append(updates, nums)
		}
			
	}

	return
}
