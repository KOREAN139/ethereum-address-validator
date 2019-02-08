package validate

import (
	"regexp"
)

func isValidFormat(addr string) bool {
	m, _ := regexp.MatchString(`^0x[a-fA-F0-9]{40}$`, addr)
	return m
}
