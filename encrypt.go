package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	key, _ := base64.StdEncoding.DecodeString("HJkPmTz+uY7wd0p1+w//DABgbvPq9/230RwEG2sJ9mo=")
	iv, _ := base64.StdEncoding.DecodeString("AAAAAAAAAAAAAAAAAAAAAA==")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := []byte(scanner.Text())
		padLen := aes.BlockSize - (len(data) % aes.BlockSize)
		padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
		paddedData := append(data, padding...)
		ct := make([]byte, len(paddedData))
		cbc := cipher.NewCBCEncrypter(block, iv)
		cbc.CryptBlocks(ct, paddedData)
		ctOut := base64.StdEncoding.EncodeToString(ct)
		fmt.Println(ctOut)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
