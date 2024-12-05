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

type Rule struct {
	Rule []int
	Respected bool
	Matched bool
}


func main() {
	//fmt.Println(part1("example_input.txt"))
	//fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}


func part2(filename string) (result int) {
	rules, updates := readFile(filename)

	for _, update := range updates {
		rulesRespected := checkUpdate(update, rules)
		if !rulesRespected {
			fixedUpdate := orderUpdate(update, rules)
			if (len(fixedUpdate)-1) % 2 == 0 {
				result = result + fixedUpdate[(len(fixedUpdate)-1) / 2]
			}
		}
	}


	return
}


func orderUpdate(update []int, rules []Rule) (fixedUpdate []int) {
	fixedUpdate = make([]int, len(update))
	copy(fixedUpdate, update)

	//for !checkUpdate(fixedUpdate, rules) {
		sort.Slice(fixedUpdate, func(i, j int) bool {
			a := fixedUpdate[i]
			b := fixedUpdate[j]
			for _, rule := range rules {
				if slices.Contains(rule.Rule, a) && slices.Contains(rule.Rule, b) {
					if numIsBefore(a, rule) {
						return slices.Index(update, a) < slices.Index(update, b)
					} else if numIsAfter(a, rule) {
						return slices.Index(update, a) > slices.Index(update, b)
					}
				}
			}
			return false
		})
	//}


	return
}


func numIsBefore(num int, rule Rule) bool {
	if rule.Rule[0] == num {
		return true
	}
	return false
}

func numIsAfter(num int, rule Rule) bool {
	if rule.Rule[1] == num {
		return true
	}
	return false
}

func part1(filename string) (result int) {
	rules, updates := readFile(filename)

	for _, update := range updates {
		rulesRespected := checkUpdate(update, rules)
		if rulesRespected && (len(update)-1) % 2 == 0 {
			result = result + update[(len(update)-1) / 2]
		}
	}

	return
}


func checkUpdate(update []int, rules []Rule) (allRulesRespected bool) {
	var rulesMatched []Rule
	allRulesRespected = true

	for orderId := range update {
		rulesMatched = checkRules(update, orderId, rules)
		for _, ruleMatched := range rulesMatched {
			if !ruleMatched.Respected {
				allRulesRespected = false
			}
		}
	}

	return allRulesRespected
}


// Checks all rules for a given number in the update, referred to by orderId
func checkRules(update []int, orderId int, rules []Rule) (rulesMatched []Rule) {
	num := update[orderId]

	for _, rule := range rules {
		if numIsBefore(num, rule) && slices.Contains(update, rule.Rule[1]) {
			matchedRule := Rule{
				Rule: rule.Rule,
				Respected: false,
				Matched: true,
			}
			if slices.Index(update, rule.Rule[0]) < slices.Index(update, rule.Rule[1]) {
				matchedRule.Respected = true
			} 
			rulesMatched = append(rulesMatched, matchedRule)
		}
		if numIsAfter(num, rule) && slices.Contains(update, rule.Rule[0]) {
			matchedRule := Rule{
				Rule: rule.Rule,
				Respected: false,
				Matched: true,
			}
			if slices.Index(update, rule.Rule[0]) < slices.Index(update, rule.Rule[1]) {
				matchedRule.Respected = true
			}
			rulesMatched = append(rulesMatched, matchedRule)
		}
	}

	return rulesMatched
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
