package crack

import (
	"crypto/sha512"
	"fmt"
)

func SHA512String(s string) string {
	sum := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", sum)
}
