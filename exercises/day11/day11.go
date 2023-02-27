package day11

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct{}

const (
	Old = iota
	Number
)

type Operation struct {
	opType1   int
	argument1 int
	operator  rune
	opType2   int
	argument2 int
}

func NewOperation(left, right string, operator rune) Operation {
	var opType1, argument1, opType2, argument2 int
	if left == "old" {
		opType1 = Old
	} else {
		opType1 = Number
		argument1, _ = strconv.Atoi(left)
	}
	if right == "old" {
		opType2 = Old
	} else {
		opType2 = Number
		argument2, _ = strconv.Atoi(right)
	}
	return Operation{opType1, argument1, operator, opType2, argument2}
}

func (o Operation) Execute(old int) int {
	var left, right int
	if o.opType1 == Old {
		left = old
	} else {
		left = o.argument1
	}
	if o.opType2 == Old {
		right = old
	} else {
		right = o.argument2
	}

	switch o.operator {
	case '+':
		return left + right
	case '*':
		return left * right
	case '-':
		return left - right
	case '/':
		return left / right
	default:
		panic("Invalid Operator " + string(o.operator))
	}
}

func (o Operation) String() string {
	var left, right string
	if o.opType1 == Old {
		left = "old"
	} else {
		left = strconv.Itoa(o.argument1)
	}
	if o.opType2 == Old {
		right = "old"
	} else {
		right = strconv.Itoa(o.argument2)
	}
	return fmt.Sprintf("%s %c %s", left, o.operator, right)
}

type Monkey struct {
	items             []int
	operation         Operation
	divisibilityTest  int
	trueTarget        int
	falseTarget       int
	inspectionCounter int
}

func ParseMonkey(str string) *Monkey {
	lines := strings.Split(str, "\n")

	itemsStr := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
	items := make([]int, len(itemsStr))
	for i, number := range itemsStr {
		items[i], _ = strconv.Atoi(number)
	}

	var operand1, operand2 string
	var operator rune
	_, _ = fmt.Sscanf(lines[2], "  Operation: new = %s %c %s", &operand1, &operator, &operand2)
	operation := NewOperation(operand1, operand2, operator)

	var divisible int
	_, _ = fmt.Sscanf(lines[3], "  Test: divisible by %d", &divisible)

	var monkeyTrue int
	_, _ = fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &monkeyTrue)
	var monkeyFalse int
	_, _ = fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &monkeyFalse)

	return &Monkey{
		items:            items,
		operation:        operation,
		divisibilityTest: divisible,
		trueTarget:       monkeyTrue,
		falseTarget:      monkeyFalse,
	}
}

func (m *Monkey) AddItem(worry, modulo int) {
	m.items = append(m.items, worry%modulo)
}

func (m *Monkey) Step(divide bool) (monkey, item int, hasNextStep bool, err error) {
	if len(m.items) == 0 {
		return 0, 0, false, errors.New("no items")
	}
	item, m.items = m.items[0], m.items[1:]
	item = m.operation.Execute(item)
	if divide {
		item /= 3
	}
	if item%m.divisibilityTest == 0 {
		monkey = m.trueTarget
	} else {
		monkey = m.falseTarget
	}
	m.inspectionCounter++
	return monkey, item, len(m.items) > 0, nil
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / (gcd(a, b))
}

func (s Solver) SolvePart1(input string) string {
	split := strings.Split(input, "\n\n")
	monkeys := make([]*Monkey, len(split))
	for i, monkeyStr := range split {
		monkeys[i] = ParseMonkey(monkeyStr)
	}

	monkeyLcm := monkeys[0].divisibilityTest
	for _, monkey := range monkeys[1:] {
		monkeyLcm = lcm(monkeyLcm, monkey.divisibilityTest)
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			hasNext := true
			var newMonkey, worry int
			var err error
			for hasNext {
				newMonkey, worry, hasNext, err = monkey.Step(true)
				if err == nil {
					monkeys[newMonkey].AddItem(worry, monkeyLcm)
				}
			}
		}
	}

	var max1, max2 int
	for _, monkey := range monkeys {
		if monkey.inspectionCounter > max1 {
			max2 = max1
			max1 = monkey.inspectionCounter
		} else if monkey.inspectionCounter > max2 {
			max2 = monkey.inspectionCounter
		}
	}

	return strconv.Itoa(max1 * max2)
}

func (s Solver) SolvePart2(input string) string {
	split := strings.Split(input, "\n\n")
	monkeys := make([]*Monkey, len(split))
	for i, monkeyStr := range split {
		monkeys[i] = ParseMonkey(monkeyStr)
	}

	monkeyLcm := monkeys[0].divisibilityTest
	for _, monkey := range monkeys[1:] {
		monkeyLcm = lcm(monkeyLcm, monkey.divisibilityTest)
	}

	for i := 0; i < 10000; i++ {
		for _, monkey := range monkeys {
			hasNext := true
			var newMonkey, worry int
			var err error
			for hasNext {
				newMonkey, worry, hasNext, err = monkey.Step(false)
				if err == nil {
					monkeys[newMonkey].AddItem(worry, monkeyLcm)
				}
			}
		}
	}

	var max1, max2 int
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times.\n", i, monkey.inspectionCounter)
		if monkey.inspectionCounter > max1 {
			max2 = max1
			max1 = monkey.inspectionCounter
		} else if monkey.inspectionCounter > max2 {
			max2 = monkey.inspectionCounter
		}
	}

	return strconv.Itoa(max1 * max2)
}
