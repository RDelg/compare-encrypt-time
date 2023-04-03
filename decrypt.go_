package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	key, _ := base64.StdEncoding.DecodeString("HJkPmTz+uY7wd0p1+w//DABgbvPq9/230RwEG2sJ9mo=")
	iv, _ := base64.StdEncoding.DecodeString("AAAAAAAAAAAAAAAAAAAAAA==")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ct, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			panic(err)
		}

		pt := make([]byte, len(ct))
		block, err := aes.NewCipher(key)
		if err != nil {
			panic(err)
		}
		cbc := cipher.NewCBCDecrypter(block, iv)
		cbc.CryptBlocks(pt, ct)

		// Unpad plaintext
		padLen := int(pt[len(pt)-1])
		pt = pt[:len(pt)-padLen]

		fmt.Println(string(pt))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
