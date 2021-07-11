package custom_errors

import (
	"fmt"

	"github.com/pkg/errors"
)

type InternalServerError struct {
	Err        error
	StatusCode int
}

func (e InternalServerError) Error() string {
	return fmt.Sprint("status code: ", e.StatusCode, ".\t", e.Err)
}

func NewInternalServerError(err error, stackTraceMsg string) InternalServerError {
	err = errors.Wrap(err, "stackTrace: "+stackTraceMsg+".\tGo")
	return InternalServerError{Err: err, StatusCode: 500}
}
