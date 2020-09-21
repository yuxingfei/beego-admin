package encrypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/astaxie/beego"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

//加密
func Encrypt(byteData []byte, aeskey []byte) string {
	xpass, err := AesEncrypt(byteData, aeskey)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(xpass)
}

//解密
func Decrypt(encryptData string, aeskey []byte) string {
	bytesPass, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		beego.Error(err)
		return ""
	}

	decryptData, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		beego.Error(err)
		return ""
	}
	return string(decryptData)
}
