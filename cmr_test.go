package cmr

import (
	"testing"

	"github.com/go-zoox/testify"
)

func TestSuccess(t *testing.T) {
	body := Success("result")

	testify.Equal(t, "200", body.Code)
	testify.Equal(t, "OK", body.Message)
	testify.Equal(t, "result", body.Result.(string))
}

func TestFail(t *testing.T) {
	Register("400001", "Business version error")

	body := Fail("400001", "result")

	testify.Equal(t, "400001", body.Code)
	testify.Equal(t, "Business version error", body.Message)
	testify.Equal(t, "result", body.Result.(string))
}

func TestSuccessCustomCodeMessage(t *testing.T) {
	body := Success("result", "201", "Business OK")

	testify.Equal(t, "201", body.Code)
	testify.Equal(t, "Business OK", body.Message)
	testify.Equal(t, "result", body.Result.(string))
}
