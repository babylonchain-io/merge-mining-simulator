package common

import (
	"fmt"
)

type funcError struct {
	Func        string
	Description string
}

func (e *funcError) Error() string {
	if e.Func != "" {
		return fmt.Sprintf("%v: %v", e.Func, e.Description)
	}
	return e.Description
}

func FuncError(f string, desc string) error {
	return &funcError{Func: f, Description: desc}
}
