package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	Rule []int
	Respected bool
	Matched bool
}


func main() {
	//fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

func part1(filename string) (result int) {
	rules, updates := readFile(filename)

	for _, update := range updates {
		if checkUpdate(update, rules) && (len(update)-1) % 2 == 0 {
			result = result + update[(len(update)-1) / 2]
		}
	}
	
	return
}


func checkUpdate(update []int, rules []Rule) (rulesViolated bool) {
	var rulesMatched []Rule

	for orderId := range update {
		rulesMatched = checkRules(update, orderId, rules)
		for _, ruleMatched := range rulesMatched {
			if !ruleMatched.Respected {
				rulesViolated = true
			}
		}

	}


	return !rulesViolated
}


// Checks all rules for a given number in the update, referred to by orderId
func checkRules(update []int, orderId int, rules []Rule) (matchedRules []Rule) {
	num := update[orderId]

	for _, rule := range rules {
		if num == rule.Rule[0] && slices.Contains(update, rule.Rule[1]) {
			matchedRule := Rule{
				Rule: rule.Rule,
				Respected: false,
				Matched: true,
			}
			if slices.Index(update, rule.Rule[0]) < slices.Index(update, rule.Rule[1]) {
				matchedRule.Respected = true
			} 
			matchedRules = append(matchedRules, matchedRule)
		}
		if num == rule.Rule[1] && slices.Contains(update, rule.Rule[0]) {
			matchedRule := Rule{
				Rule: rule.Rule,
				Respected: false,
				Matched: true,
			}
			if slices.Index(update, rule.Rule[0]) < slices.Index(update, rule.Rule[1]) {
				matchedRule.Respected = true
			}
			matchedRules = append(matchedRules, matchedRule)
		}
	}

	return matchedRules
}


func readFile(filename string) (rules []Rule, updates [][]int) {
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
			rules = append(rules, Rule{
				Rule: []int{before, after},
				Respected: false,
				Matched: false,
			})
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
