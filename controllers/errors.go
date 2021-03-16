package controllers

import (
	"fmt"
)

// TargetNotFoundError is returned when the Target for the SRL object cannot be found
type TargetNotFoundError struct {
	message string
}

func (e TargetNotFoundError) Error() string {
	return fmt.Sprintf("Target Not Found %s",
		e.message)
}
