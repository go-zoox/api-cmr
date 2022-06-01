package cmr

import "fmt"

var codeMap = map[string]string{}

// Body is the response body
type Body struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

// Success return a succes body
func Success[T any](result T, codemessage ...string) *Body {
	code := "200"
	message := "OK"

	if len(codemessage) > 0 {
		code = codemessage[0]
	}
	if len(codemessage) > 1 {
		message = codemessage[1]
	}

	return &Body{
		Code:    code,
		Message: message,
		Result:  result,
	}
}

// Fail return a fail body, bad request causes calculation error
func Fail[T any](code string, result T) *Body {
	message, ok := codeMap[code]
	if !ok {
		panic(fmt.Sprintf("undefined error code(%s)", code))
	}

	return &Body{
		Code:    code,
		Message: message,
		Result:  result,
	}
}

// Register registers a pair of code and message
func Register(code string, message string) {
	codeMap[code] = message
}
