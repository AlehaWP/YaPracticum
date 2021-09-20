package encription

import (
	"crypto/aes"
	"crypto/rand"
	"strconv"
)

var keySize int = 16

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

func EncriptInt(s int) (string, error) {
	key, err := generateRandom(keySize)
	if err != nil {
		return "", err
	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	res := make([]byte, aesblock.BlockSize())
	aesblock.Encrypt(res, []byte(strconv.Itoa(31415926)))

	return "", nil

}
