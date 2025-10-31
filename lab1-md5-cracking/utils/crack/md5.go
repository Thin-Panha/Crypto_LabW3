package crack

import (
	"crypto/md5"
	"fmt"
)

func MD5String(s string) string {
	sum := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", sum)
}
