package ethav

// Validate validates given address
func Validate(addr string) error {
	if !isValidFormat(addr) {
		return ErrorTypeFormat
	}

	if !validChecksum(addr) {
		return ErrorTypeChecksum
	}

	return nil
}
