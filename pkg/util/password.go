package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

var aesKeyForLogin = []byte(`V!7e@gaS^Y#KSRvc`)

// GeneratePassword 加密密码
func GeneratePassword(password string) (string, error) {
	pwdBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pwdBytes), nil
}

// ComparePassword 对比加密后的密码用户密码原始密码
func ComparePassword(sysPassword, userRealPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(sysPassword), []byte(userRealPassword)); err != nil {
		return false
	}
	return true
}

// AesEncrypt aes 加密
func AesEncrypt(text string) (string, error) {
	pkcs5Padding := func(ciphertext []byte, blockSize int) []byte {
		padding := blockSize - len(ciphertext)%blockSize
		padtext := bytes.Repeat([]byte{byte(padding)}, padding)
		return append(ciphertext, padtext...)
	}

	block, err := aes.NewCipher(aesKeyForLogin)
	if err != nil {
		return "", err
	}
	origData := []byte(text)
	blockSize := block.BlockSize()
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, aesKeyForLogin[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

// AesDecrypt aes解密
func AesDecrypt(crypted string) (string, error) {
	pkcs5UnPadding := func(origData []byte) []byte {
		length := len(origData)
		var unpadding int
		if length == 0 {
			unpadding = 0
		} else {
			unpadding = int(origData[length-1])
		}
		return origData[:(length - unpadding)]
	}

	block, err := aes.NewCipher(aesKeyForLogin)
	if err != nil {
		return "", err
	}
	cryptedOri, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, aesKeyForLogin[:block.BlockSize()])
	origData := make([]byte, len(cryptedOri))
	blockMode.CryptBlocks(origData, cryptedOri)
	origData = pkcs5UnPadding(origData)
	return string(origData), nil
}
