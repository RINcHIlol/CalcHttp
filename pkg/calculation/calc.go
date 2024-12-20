package calculation

import (
	"errors"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	numbers := []float64{}
	operators := []rune{}

	if len(expression) == 0 {
		return 0, errors.New("Internal server error")
	}

	i := 0
	for i < len(expression) {
		char := rune(expression[i])
		if unicode.IsDigit(char) || char == '.' {
			j := i
			for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
				j++
			}
			number, err := strconv.ParseFloat(expression[i:j], 64)
			if err != nil {
				return 0, errors.New("Expression is not valid")
			}
			numbers = append(numbers, number)
			i = j - 1
		} else if char == '(' {
			operators = append(operators, char)
		} else if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				err := applyOperation(&numbers, &operators)
				if err != nil {
					return 0, errors.New("Expression is not valid")
				}
			}
			if len(operators) == 0 || operators[len(operators)-1] != '(' {
				return 0, errors.New("Expression is not valid")
			}
			operators = operators[:len(operators)-1]
		} else if isOperator(char) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(char) {
				err := applyOperation(&numbers, &operators)
				if err != nil {
					return 0, errors.New("Expression is not valid")
				}
			}
			operators = append(operators, char)
		} else if !unicode.IsSpace(char) {
			return 0, errors.New("Expression is not valid")
		}
		i++
	}
	for len(operators) > 0 {
		err := applyOperation(&numbers, &operators)
		if err != nil {
			return 0, errors.New("Internal server error")
		}
	}
	if len(numbers) != 1 {
		return 0, errors.New("Internal server error")
	}

	return numbers[0], nil
}

func applyOperation(numbers *[]float64, operators *[]rune) error {
	if len(*numbers) < 2 || len(*operators) == 0 {
		return errors.New("bello")
	}
	b := (*numbers)[len(*numbers)-1]
	a := (*numbers)[len(*numbers)-2]
	*numbers = (*numbers)[:len(*numbers)-2]
	op := (*operators)[len(*operators)-1]
	*operators = (*operators)[:len(*operators)-1]

	var result float64
	switch op {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			return errors.New("bello")
		}
		result = a / b
	default:
		return errors.New("bello")
	}
	*numbers = append(*numbers, result)
	return nil
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	default:
		return 0
	}
}
