package utils

import (
	"log"
	"regexp"
	"strings"
)

func ValidateExpression(expression string) (bool, string) {
	if !strings.HasPrefix(expression, "What is") || !strings.HasSuffix(expression, "?") {
		return false, "Expression should start with 'What is' and end with '?'"
	}

	expression = strings.TrimPrefix(expression, "What is ")
	tokens := strings.Fields(strings.TrimSpace(expression))

	isValidNumber := regexp.MustCompile(`^\d+$`).MatchString
	isValidLastNumber := regexp.MustCompile(`^\d+\?$`).MatchString
	log.Printf("Expression so far: %s", expression)

	expected := "number"

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		switch expected {
		case "number":
			log.Printf("In number case: %s", token)
			if isValidLastNumber(token) {
				log.Printf("In valid last num case: %s", token)
				return true, ""
			}
			if !isValidNumber(token) {
				return false, "Invalid syntax: expected a number, got '" + token + "'"
			}
			expected = "operator"

		case "operator":
			log.Printf("In operator case: %s", token)
			if token == "plus" || token == "minus" {
				expected = "number"
			} else if token == "multiplied" || token == "divided" {
				if i+1 >= len(tokens) || tokens[i+1] != "by" {
					return false, "Invalid syntax: expected 'by' after '" + token + "'"
				}
				i++
				expected = "number"
			} else {
				if isValidNumber(token) {
					return false, "Expected operation, received number: '" + token + "' instead."
				}
				return false, "Unsupported operation: '" + token + "'"
			}
		}
	}

	if expected == "operator" {
		log.Printf("In fail case: %s", expected)
		return false, "Expression ends with an operator, expected a number."
	}

	return true, ""
}
