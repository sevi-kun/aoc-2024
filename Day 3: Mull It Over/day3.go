package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	var result int

	re_muls := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	matched_muls := re_muls.FindAll([]byte(data), -1)
	
	re_num := regexp.MustCompile(`[0-9]{1,3}`)

	for _, mul := range matched_muls {
		numbers_dirty := strings.Split(string(mul), ",")
		num1, err1 := strconv.Atoi(
			string(
				re_num.Find(
					[]byte(numbers_dirty[0]))))

		num2, err2 := strconv.Atoi(string(re_num.Find(
					[]byte(numbers_dirty[1]))))

		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing numbers: ", numbers_dirty)
		}

		result = result + num1 * num2
	}



	return result
}


func readFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = data + scanner.Text()
	}

	return data
}
