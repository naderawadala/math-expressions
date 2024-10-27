package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func CalculateExpression(expression string) (int, error) {
	expression = strings.TrimPrefix(expression, "What is ")
	tokens := strings.Fields(strings.TrimSpace(expression))

	isValidNumber := regexp.MustCompile(`^\d+$`).MatchString
	isValidLastNumber := regexp.MustCompile(`^\d+\?$`).MatchString

	result := 0
	expected := "number"
	var operator string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch expected {
		case "number":
			if isValidLastNumber(token) {
				token = strings.TrimSuffix(token, "?")
				num, err := strconv.Atoi(token)
				if err != nil {
					return 0, fmt.Errorf("failed to parse number: %s", token)
				}
				result = calculateOperation(result, num, operator)
				return result, nil
			}
			if isValidNumber(token) {
				num, err := strconv.Atoi(token)
				if err != nil {
					return 0, fmt.Errorf("failed to parse number: %s", token)
				}

				result = calculateOperation(result, num, operator)
				expected = "operator"
			} else {
				return 0, fmt.Errorf("invalid number: %s", token)
			}

		case "operator":
			switch token {
			case "plus", "minus":
				operator = token
				expected = "number"
			case "multiplied", "divided":
				if i+1 >= len(tokens) || tokens[i+1] != "by" {
					return 0, fmt.Errorf("invalid syntax: expected 'by' after '%s' ", token)
				}
				i++
				operator = token
				expected = "number"
			default:
				return 0, fmt.Errorf("unsupported operation: %s", token)
			}
		}
	}

	return result, errors.New("incomplete expression")
}

func calculateOperation(result int, num int, operator string) int {
	switch operator {
	case "":
		result = num
	case "plus":
		result += num
	case "minus":
		result -= num
	case "multiplied":
		result *= num
	case "divided":
		if num == 0 {
			return 0
		}
		result /= num
	}
	return result
}
