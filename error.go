package try

import "errors"

// ErrMaxRetriesReached exceeded retry limit
var ErrMaxRetriesReached = errors.New("exceeded retry limit")
var ErrSomethingWrong = errors.New("something wrong")
