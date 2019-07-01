package main

import (
	"fmt"

	"github.com/oyakata/gooooolang/seckeyenc"
)

func main() {
	baseValue := "1234-abcd-5678-efgh"

	value, err := seckeyenc.Encrypt([]byte(baseValue))
	fmt.Println(string(value), err)

	result, err := seckeyenc.Decrypt(value)
	fmt.Println(string(result), err)

	s, err := seckeyenc.EncryptBase64URLString([]byte(baseValue))
	fmt.Println("Base64URL:", s, err)

	b, err := seckeyenc.DecryptBase64URLString(s)
	fmt.Println("Base64URL:", string(b), err)
}
