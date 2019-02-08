package validate

import (
	"golang.org/x/crypto/sha3"
	"strings"
)

func validChecksum(addr string) bool {
	hex := strings.ToLower(addr)[2:]

	d := sha3.NewLegacyKeccak256()
	d.Write([]byte(hex))
	hash := d.Sum(nil)

	ret := "0x"

	for i, b := range hex {
		c := string(b)
		if b < '0' || b > '9' {
			curr := hash[i/2] >> uint(4-i%2*4)
			if curr&8 != 0 {
				c = string(b - 32)
			}
		}
		ret += c
	}

	return addr == ret
}
