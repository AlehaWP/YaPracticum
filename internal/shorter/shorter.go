package shorter

import (
	"crypto/md5"
)

//MD5 func for hash.
func MD5(b []byte) string {
	h := md5.Sum(b)
	return fmt.Sprintf("%x", h)
}
