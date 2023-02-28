package day13

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

type Solver struct{}

const (
	InOrder = iota
	NotInOrder
	Inconclusive
)

var OrderNames = map[int]string{
	InOrder:      "In Order",
	NotInOrder:   "Not In Order",
	Inconclusive: "Order Inconclusive",
}

func Compare(a, b []interface{}) int {
	for i := 0; ; i++ {
		if i >= len(a) && i >= len(b) {
			return Inconclusive
		} else if i >= len(a) {
			return InOrder
		} else if i >= len(b) {
			return NotInOrder
		}

		var aInt, bInt bool
		switch a[i].(type) {
		case float64:
			aInt = true
		}
		switch b[i].(type) {
		case float64:
			bInt = true
		}
		if aInt && bInt {
			if a[i].(float64) < b[i].(float64) {
				return InOrder
			} else if a[i].(float64) > b[i].(float64) {
				return NotInOrder
			}
		} else if !aInt && !bInt {
			if order := Compare(a[i].([]interface{}), b[i].([]interface{})); order != Inconclusive {
				return order
			}
		} else if aInt {
			if order := Compare([]interface{}{a[i].(float64)}, b[i].([]interface{})); order != Inconclusive {
				return order
			}
		} else {
			if order := Compare(a[i].([]interface{}), []interface{}{b[i].(float64)}); order != Inconclusive {
				return order
			}
		}
	}
}

func (s Solver) SolvePart1(input string) string {
	var sum int
	for i, pair := range strings.Split(input, "\n\n") {
		split := strings.Split(pair, "\n")
		var a, b []interface{}
		_ = json.Unmarshal([]byte(split[0]), &a)
		_ = json.Unmarshal([]byte(split[1]), &b)
		if Compare(a, b) == InOrder {
			sum += i + 1
		}
	}

	return strconv.Itoa(sum)
}

func (s Solver) SolvePart2(input string) string {
	var lists []interface{}

	input = input + "\n\n[[2]]\n[[6]]"

	for _, pair := range strings.Split(input, "\n\n") {
		split := strings.Split(pair, "\n")
		var a, b []interface{}
		_ = json.Unmarshal([]byte(split[0]), &a)
		_ = json.Unmarshal([]byte(split[1]), &b)
		lists = append(append(lists, a), b)
	}

	sort.Slice(lists, func(i, j int) bool {
		return Compare(lists[i].([]interface{}), lists[j].([]interface{})) == InOrder
	})

	var first, second int
	for i := 0; i < len(lists); i++ {
		switch lists[i].(type) {
		case []interface{}:
			list := lists[i].([]interface{})
			if len(list) == 1 {
				switch list[0].(type) {
				case []interface{}:
					if len(list) == 1 {
						list = list[0].([]interface{})
						if len(list) == 1 {
							switch list[0].(type) {
							case float64:
								element := list[0].(float64)
								if element == 2 {
									first = i + 1
								} else if element == 6 {
									second = i + 1
								}
							}
						}
					}
				}
			}
		}
	}

	return strconv.Itoa(first * second)
}
