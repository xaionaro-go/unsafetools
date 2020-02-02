package unsafetools

import (
	"fmt"
)

type ErrSliceTooShort struct {
	RequiredSize uintptr
	RealSize     uintptr
}

func (err *ErrSliceTooShort) Error() string {
	return fmt.Sprintf("the slice is too short: %v < %v",
		err.RealSize, err.RequiredSize)
}
