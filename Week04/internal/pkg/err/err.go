package err

import "github.com/pkg/errors"

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrDecodeJsonError = errors.New("decode json error")
	ErrNotFoundCommand = errors.New("not found command")
	ErrNeedParam       = errors.New("need command param")
	ErrParamTypeError  = errors.New("param type error")
)
