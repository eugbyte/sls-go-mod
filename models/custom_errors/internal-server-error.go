package custom_errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type InternalServerError struct {
	Err        error
	StatusCode int
}

func (e *InternalServerError) Error() string {
	return fmt.Sprint("status code: ", e.StatusCode, ".\t", e.Err)
}

func NewInternalServerError(err error, stackTraceMsg string) BadRequest {
	err = errors.Wrap(err, "stackTrace: "+stackTraceMsg+".\tGo")
	return BadRequest{Err: err, StatusCode: 500}
}
