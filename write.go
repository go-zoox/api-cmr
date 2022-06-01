package cmr

import (
	"net/http"

	"github.com/go-zoox/encoding/json"
)

// WriteSuccess writes a success response to the response writer.
func WriteSuccess(w http.ResponseWriter, result interface{}, codemessage ...string) {
	body := Success(result, codemessage...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, _ := json.Encode(body)

	w.Write(bytes)
}

// WriteFail writes a fail response to the response writer.
func WriteFail(w http.ResponseWriter, code string, message string, result interface{}) {
	body := Fail(code, result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	bytes, _ := json.Encode(body)

	w.Write(bytes)
}
