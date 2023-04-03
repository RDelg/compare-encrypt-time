package services

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type CypherService interface {
	Encrypt(message string) (string, error)
	Decrypt(message string) (string, error)
}

type cypherService struct {
	key   []uint8
	iv    []uint8
	block cipher.Block
}

func NewCypherService(key string, iv string) CypherService {
	key_bin, _ := base64.StdEncoding.DecodeString(key)
	iv_bin, _ := base64.StdEncoding.DecodeString(iv)
	block, err := aes.NewCipher(key_bin)
	if err != nil {
		panic(err)
	}
	return &cypherService{key: key_bin, iv: iv_bin, block: block}
}

func (s *cypherService) Encrypt(data string) (string, error) {
	// process data
	data_bin := []byte(data)
	padLen := aes.BlockSize - (len(data_bin) % aes.BlockSize)
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	paddedData := append(data_bin, padding...)
	// encrypt data
	cbc := cipher.NewCBCEncrypter(s.block, s.iv)
	ct := make([]byte, len(paddedData))
	cbc.CryptBlocks(ct, paddedData)
	ctOut := base64.StdEncoding.EncodeToString(ct)
	return ctOut, nil
}

func (s *cypherService) Decrypt(message string) (string, error) {
	ct, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		panic(err)
	}
	pt := make([]byte, len(ct))
	cbc := cipher.NewCBCDecrypter(s.block, s.iv)
	cbc.CryptBlocks(pt, ct)

	// Unpad plaintext
	padLen := int(pt[len(pt)-1])
	pt = pt[:len(pt)-padLen]
	return string(pt), nil
}
