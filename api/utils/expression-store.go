package utils

type Errors struct {
	Expression string `json:"expression"`
	Endpoint   string `json:"endpoint"`
	Frequency  int    `json:"frequency"`
	ErrorType  string `json:"type"`
}

var errorLog = map[string]*Errors{}

func LogError(expression, endpoint, errorType string) {
	if record, exists := errorLog[expression]; exists {
		record.Frequency++
	} else {
		errorLog[expression] = &Errors{
			Expression: expression,
			Endpoint:   endpoint,
			Frequency:  1,
			ErrorType:  errorType,
		}
	}
}

func GetErrors() []Errors {
	errors := make([]Errors, 0, len(errorLog))
	for _, record := range errorLog {
		errors = append(errors, *record)
	}
	return errors
}
