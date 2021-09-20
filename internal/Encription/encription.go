package encription

import (
	"crypto/aes"
	"crypto/md5"
	"crypto/rand"
	"fmt"
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

func EncriptInt(i int) (string, error) {
	key, err := generateRandom(aes.BlockSize)
	fmt.Println("keySize : ", keySize)
	fmt.Println("key : ", key)
	if err != nil {
		return "", err
	}

	aesblock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	fmt.Println("BlockSize : ", aesblock.BlockSize())

	res := make([]byte, aesblock.BlockSize())
	fmt.Println("i : ", i)
	fmt.Println("res : ", res)
	hash := md5.Sum([]byte(strconv.Itoa(i)))
	aesblock.Encrypt(res, hash[:])
	fmt.Printf("encrypted: %q\n", res)

	return fmt.Sprintf("%x", res), nil

}
