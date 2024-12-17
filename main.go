package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"unicode"
)

type Request struct {
	Expression string `json:"expression"`
}

type Result struct {
	Result     float64 `json:"result"`
	StatusCode int     `json:"statusCode"`
}

type ErrorResult struct {
	Error      string `json:"error"`
	StatusCode int    `json:"statusCode"`
}

func main() {
	http.HandleFunc("/", Calc)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("error4ik")
	}
}

func Calc(w http.ResponseWriter, r *http.Request) {
	numbers := []float64{}
	operators := []rune{}
	var request Request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		errorData := ErrorResult{Error: "Expression is not valid", StatusCode: http.StatusUnprocessableEntity}
		errorJson, _ := json.Marshal(errorData)
		w.Write(errorJson)
		return
	}
	expression := request.Expression

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
				errorData := ErrorResult{Error: "Expression is not valid", StatusCode: http.StatusUnprocessableEntity}
				errorJson, _ := json.Marshal(errorData)
				w.Write(errorJson)
				return
			}
			numbers = append(numbers, number)
			i = j - 1
		} else if char == '(' {
			operators = append(operators, char)
		} else if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				err := applyOperation(&numbers, &operators)
				if err != nil {
					errorData := ErrorResult{Error: "Expression is not valid", StatusCode: http.StatusUnprocessableEntity}
					errorJson, _ := json.Marshal(errorData)
					w.Write(errorJson)
					return
				}
			}
			if len(operators) == 0 || operators[len(operators)-1] != '(' {
				errorData := ErrorResult{Error: "Expression is not valid", StatusCode: http.StatusUnprocessableEntity}
				errorJson, _ := json.Marshal(errorData)
				w.Write(errorJson)
				return
			}
			operators = operators[:len(operators)-1]
		} else if isOperator(char) {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(char) {
				err := applyOperation(&numbers, &operators)
				if err != nil {
					errorData := ErrorResult{Error: "Expression is not valid", StatusCode: http.StatusUnprocessableEntity}
					errorJson, _ := json.Marshal(errorData)
					w.Write(errorJson)
					return
				}
			}
			operators = append(operators, char)
		} else if !unicode.IsSpace(char) {
			errorData := ErrorResult{Error: "Expression is not valid", StatusCode: http.StatusUnprocessableEntity}
			errorJson, _ := json.Marshal(errorData)
			w.Write(errorJson)
			return
		}
		i++
	}
	for len(operators) > 0 {
		err := applyOperation(&numbers, &operators)
		if err != nil {
			errorData := ErrorResult{Error: "Internal server error", StatusCode: http.StatusInternalServerError}
			errorJson, _ := json.Marshal(errorData)
			w.Write(errorJson)
			return
		}
	}
	if len(numbers) != 1 {
		errorData := ErrorResult{Error: "Internal server error", StatusCode: http.StatusInternalServerError}
		errorJson, _ := json.Marshal(errorData)
		w.Write(errorJson)
		return
	}

	response := Result{Result: numbers[0], StatusCode: http.StatusOK}

	responseJson, err := json.Marshal(response)
	if err != nil {
		errorData := ErrorResult{Error: "Internal server error", StatusCode: http.StatusInternalServerError}
		errorJson, _ := json.Marshal(errorData)
		w.Write(errorJson)
		return
	}

	w.Write(responseJson)
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
