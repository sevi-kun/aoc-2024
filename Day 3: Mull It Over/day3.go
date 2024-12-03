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
	//fmt.Println(part1("example_input1.txt"))
	//fmt.Println(part1("input.txt"))
	//mt.Println(part2("example_input2.txt"))
	fmt.Println(part2("input.txt"))
}


func part2(filename string) int {
	data := readFile(filename)
	var result int
	var mul_count int

	re_do := regexp.MustCompile(`do\(\)`)
	do_index_list := re_do.FindAllStringIndex(data, -1)
	fmt.Println(len(do_index_list))

	do_index_pre := 0
	for _, do_index := range do_index_list {

		mul_string := strings.Split(
				data[do_index_pre:do_index[0]], 
				"don't()")[0]

		mulreturn := matchMulMul(mul_string)
		
		result = result + mulreturn[0]
		mul_count = mul_count + mulreturn[1]

		do_index_pre = do_index[1]
	}

	mul_string := strings.Split(
		data[do_index_pre:len(data)-1], 
		"don't()")[0]

	mulreturn := matchMulMul(mul_string)
	
	result = result + mulreturn[0]
	mul_count = mul_count + mulreturn[1]

	fmt.Println("Mul count: ", mul_count)
	return result
}


func part1(filename string) int {
	data := readFile(filename)
	return matchMulMul(data)[0]
}

// matches muls and multiplies them
func matchMulMul(data string) []int  {
	var result int
	var mul_count int

	re_muls := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	re_num := regexp.MustCompile(`[0-9]{1,3}`)
	mach_muls := re_muls.FindAllString(data, -1)
	mul_count = len(mach_muls)

	for _, mul := range mach_muls {
		numbers_dirty := strings.Split(mul, ",")
		num1, err1 := strconv.Atoi(
				re_num.FindString(
					numbers_dirty[0]))

		num2, err2 := strconv.Atoi(re_num.FindString(
			numbers_dirty[1]))

		if err1 != nil || err2 != nil {
			log.Fatal("Error parsing numbers: ", numbers_dirty)
		}

		result = result + num1 * num2
	}

	return []int{result, mul_count}

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
