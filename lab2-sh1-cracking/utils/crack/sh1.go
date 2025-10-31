package crack

import (
	"crypto/sha1"
	"fmt"
)

func SHA1String(s string) string {
	sum := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", sum)
}
