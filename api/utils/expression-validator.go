package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateExpression(expression string) (bool, error) {
	if !strings.HasPrefix(expression, "What is") || !strings.HasSuffix(expression, "?") {
		return false, errors.New("expression should start with 'What is' and end with '?'")
	}

	expression = strings.TrimPrefix(expression, "What is ")
	tokens := strings.Fields(strings.TrimSpace(expression))

	isValidNumber := regexp.MustCompile(`^\d+$`).MatchString
	isValidLastNumber := regexp.MustCompile(`^\d+\?$`).MatchString

	expected := "number"

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch expected {
		case "number":
			if isValidLastNumber(token) {
				return true, nil
			}
			if !isValidNumber(token) {
				return false, errors.New("invalid syntax: expected a number, got '" + token + "'")
			}
			expected = "operator"

		case "operator":
			if token == "plus" || token == "minus" {
				expected = "number"
			} else if token == "multiplied" || token == "divided" {
				if i+1 >= len(tokens) || tokens[i+1] != "by" {
					return false, errors.New("invalid syntax: expected 'by' after '" + token + "'")
				}
				i++
				expected = "number"
			} else {
				if isValidNumber(token) {
					return false, errors.New("expected operation, received number: '" + token + "' instead")
				}
				return false, errors.New("unsupported operation: '" + token + "'")
			}
		}
	}

	if expected == "operator" {
		return false, errors.New("expression ends with an operator, expected a number")
	}

	return true, nil
}
