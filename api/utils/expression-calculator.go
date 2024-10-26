package utils

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func CalculateExpression(expression string) (int, error) {
	expression = strings.TrimPrefix(expression, "What is ")
	tokens := strings.Fields(strings.TrimSpace(expression))

	isValidNumber := regexp.MustCompile(`^\d+$`).MatchString
	isValidLastNumber := regexp.MustCompile(`^\d+\?$`).MatchString

	log.Printf("Expression so far: %s", expression)

	result := 0
	expected := "number"
	var operator string

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch expected {
		case "number":
			if isValidLastNumber(token) {
				log.Printf("Expression so far in IS VALID LAST NUMBER: %s", expression)
				token = strings.TrimSuffix(token, "?")
				log.Printf("TOKEN IN IS VALID LAST NUM: %s", token)
				num, err := strconv.Atoi(token)
				if err != nil {
					return 0, fmt.Errorf("failed to parse number: %s", token)
				}
				result = calculateOperation(result, num, operator)
				log.Printf("RETURN RSULT IN IS VALID LAST NUM: %d", result)
				return result, nil
			}
			if isValidNumber(token) {
				log.Printf("TOKEN IN NUM: %s", token)
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
				log.Printf("OPEARTOR HAS BEEN SET TO: %s", token)
				operator = token
				expected = "number"
			case "multiplied", "divided":
				if i+1 >= len(tokens) || tokens[i+1] != "by" {
					return 0, fmt.Errorf("Invalid syntax: expected 'by' after '%s' ", token)
				}
				i++
				operator = token
				expected = "number"
				log.Printf("OPEARTOR DIVISION OR MULTIPLICATION SELECTED: %s", token)
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
		log.Printf("INITIAL SET: %d", result)
	case "plus":
		log.Printf("RESULT EQUALS :%d WITH OPERATOR %s AND NUMBER: %d", result, operator, num)
		result += num
	case "minus":
		log.Printf("RESULT EQUALS :%d WITH OPERATOR %s AND NUMBER: %d", result, operator, num)
		result -= num
	case "multiplied":
		log.Printf("RESULT EQUALS :%d WITH OPERATOR %s AND NUMBER: %d", result, operator, num)
		result *= num
	case "divided":
		log.Printf("RESULT EQUALS :%d WITH OPERATOR %s AND NUMBER: %d", result, operator, num)
		if num == 0 {
			return 0
		}
		result /= num
	}
	log.Printf("RETURN RESULT: %d", result)
	return result
}
