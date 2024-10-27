# Math-Expressions

Math-Expressions is a web service written on Go that evaluates simple math word problems. It provides endpoints for parsing and evaluating expressions, validating their syntax, and retrieving error logs of invalid evaluations.

## Installation
Clone the repository, all you need as a prerequisite is to have the Go programming language installed on your machine.

```
git clone https://github.com/naderawadala/math-expressions.git
```

In order to run the project run the following terminal command in the directory of the main.go file:

```
go run main.go
```

The server will be listening on http://localhost:8080

## Routes

### POST /evaluate

**Description**: Evaluates a math expression and returns the result of the operation if valid, otherwise returns an error with the reason why the math expression is invalid.

Valid operators are: "plus", "minus", "divided by", "multiplied by". Trying to use any other operator will return an Unsupported operation error.

**Request Body**:

```json
{
  "expression": "<a simple math problem>"
}
```

**Response Body**:

```json
{
  "result": "<result from expression operations>"
}
```

```json
{
  "error": "<error description>"
}
```

**Example:**

***Request:***

```json
{
  "expression": "What is 5 plus 13?"
}
```

***Response:***

```json
{
  "result": 18
}
```

***Request:***

```json
{
  "expression": "What is 1 plus plus 2?"
}
```

***Response:***

```json
{
  "error": "Invalid syntax: expected a number, got 'plus'"
}
```

### POST /validate

**Description**: Validates a math expression and returns whether it is a valid expression or not. If invalid also returns the reason.

**Request Body**:

```json
{
  "expression": "<a simple math problem>"
}
```

**Response Body**:

```json
{
  "valid": "<true or false>"
  "reason": "<if false, returns why the expression is invalid, otherwise this does not get returned>"
}
```

**Example:**

***Request:***

```json
{
  "expression": "What is 5 plus 13?"
}
```

***Response:***

```json
{
  "valid": true
}
```

***Request:***

```json
{
  "expression": "What is 1 plus plus 2?"
}
```

***Response:***

```json
{
  "valid": false
  "reason": "Invalid syntax: expected a number, got 'plus'"
}
```

### GET /errors

**Description**: Returns all the invalid expressions stored in memory.

**Response Body**:

```json
{
"result": [
  {
    "expression": "<a simple math problem>",
    "endpoint": "<endpoint URL where the error occured>",
    "frequency": "<number of times the expression failed>",
    "type": "<error type, i.e 'unsupported operation' or 'invalid syntax'>"
  },
  ...
]
}
```

**Example:**
```json
{
    "result": [
        {
            "expression": "What is 5 plus 4 minus ten?",
            "endpoint": "/evaluate",
            "frequency": 1,
            "type": "invalid syntax: expected a number, got 'ten?'"
        },
        {
            "expression": "What is 5 plus 4 10?",
            "endpoint": "/validate",
            "frequency": 3,
            "type": "unsupported operation: '10?'"
        },
        {
            "expression": "What is 5 plus 4 minus?",
            "endpoint": "/evaluate",
            "frequency": 1,
            "type": "unsupported operation: 'minus?'"
        }
    ]
}
```