package custom_errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type BadRequest struct {
	Err        error
	StatusCode int
}

func (e BadRequest) Error() string {
	return fmt.Sprint("status code: ", e.StatusCode, ".\t", e.Err)
}

func NewBadRequest(err error, stackTraceMsg string) BadRequest {
	err = errors.Wrap(err, "stackTrace: "+stackTraceMsg+".\tGo")
	return BadRequest{Err: err, StatusCode: 400}
}
