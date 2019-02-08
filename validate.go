package ethav

// Validate validates given address
func Validate(addr string) error {
	if !validFormat(addr) {
		return ErrorTypeFormat
	}

	if !validChecksum(addr) {
		return ErrorTypeChecksum
	}

	return nil
}
