package utils

import (
	"math"
	"os"
	"strconv"
	"strings"
)

const Plus = "+"
const Multiply = "*"
const Concat = "||"

type Equation struct {
	Result      int
	Values      []int
	Expressions []string
}

func (eq *Equation) IsCorrect() bool {
	result := eq.Values[0]
	for i := 1; i < len(eq.Values); i++ {
		if eq.Expressions[i-1] == Plus {
			result += eq.Values[i]
		} else if eq.Expressions[i-1] == Multiply {
			result *= eq.Values[i]
		} else {
			result, _ = strconv.Atoi(strconv.Itoa(result) + strconv.Itoa(eq.Values[i]))
		}
	}
	return eq.Result == result
}

func (eq *Equation) AddExpression(expression string) {
	eq.Expressions = append(eq.Expressions, expression)
}

func (eq *Equation) SetExpression(ex []string) {
	eq.Expressions = ex
}

func GetInput() []Equation {
	fileData, _ := os.ReadFile("input.txt")
	rows := strings.Split(string(fileData), "\n")
	equations := make([]Equation, len(rows))

	for i, row := range rows {
		resultSlice := strings.Split(row, ":")
		result, _ := strconv.Atoi(resultSlice[0])

		valuesSlice := strings.Split(strings.TrimLeft(resultSlice[1], " "), " ")
		values := make([]int, len(valuesSlice))

		for j, value := range valuesSlice {
			values[j], _ = strconv.Atoi(value)
		}

		equations[i] = Equation{
			Result:      result,
			Values:      values,
			Expressions: make([]string, 0),
		}
	}

	return equations
}

func CreatePermutations(equation Equation, outcomes []string) []Equation {
	permutations := make([]Equation, int(math.Pow(float64(len(outcomes)), float64(len(equation.Values)-1))))

	for i := 0; i < len(permutations); i++ {
		permutation := equation
		num := i
		for j := 0; j < len(permutation.Values)-1; j++ {
			digit := num % len(outcomes) // Get remainder (0, 1, or 2)
			num /= len(outcomes)         // Divide by 3 for next iteration

			permutation.AddExpression(outcomes[digit])
		}
		permutations[i] = permutation
	}
	return permutations
}

func Solve(outcomes []string) int {
	equations := GetInput()
	sum := 0
	for _, equation := range equations {
		for _, permutation := range CreatePermutations(equation, outcomes) {

			if permutation.IsCorrect() {
				sum += permutation.Result
				break
			}
		}
	}

	return sum
}
