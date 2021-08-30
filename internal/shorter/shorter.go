package shorter

import (
	"crypto/md5"
	"fmt"
)

//MD5 func for hash.
func MakeShortner(b []byte) string {
	h := md5.Sum(b)
	return fmt.Sprintf("%x", h)
}
