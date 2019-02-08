package validate

import (
	"errors"
)

var (
	// ErrorTypeFormat is set when address has invalid format
	ErrorTypeFormat = errors.New("Invalid format of ethereum address")
	// ErrorTypeChecksum is set when address has wrong checksum
	ErrorTypeChecksum = errors.New("Invalid checksum")
)
