package shorter

import (
	"crypto/md5"
	"fmt"
)

//MakeShortner func for  hash.
func MakeShortner(s string) string {
	if len(s) == 0 {
		return ""
	}
	h := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}
