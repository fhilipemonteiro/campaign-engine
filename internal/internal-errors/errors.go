package internal_errors

import "errors"

var ErrInternal error = errors.New("internal server error")
