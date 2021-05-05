package aesutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"

	"github.com/warrior21st/goutils/base58util"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

const iterationCount int = 1000

var defaultSalt = []byte{8, 6, 2, 4, 55, 0, 31, 79}

//使用默认设置加密
func AesEncryptByDefault(orig string, key string) string {
	return AesEncrypt(orig, key, defaultSalt, true)
}

//使用默认设置解密
func AesDecryptByDefault(orig string, key string) string {
	return AesDecrypt(orig, key, defaultSalt, true)
}

//aes加密
func AesEncrypt(orig string, key string, salt []byte, useBase58 bool) string {
	// 转成字节数组
	origData := []byte(orig)
	shab := sha256.Sum256([]byte(key))
	k := pbkdf2.Key(shab[:], salt, iterationCount, 32, sha3.New256)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = pkcs7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	if useBase58 {
		return base58util.Base58Encode(cryted)
	} else {
		return base64.StdEncoding.EncodeToString(cryted)
	}
}

//aes解密
func AesDecrypt(cryted string, key string, salt []byte, useBase58 bool) string {
	var crytedByte []byte
	// 转成字节数组
	if useBase58 {
		crytedByte = base58util.Base58DecodeToBytes([]byte(cryted))
	} else {
		crytedByte, _ = base64.StdEncoding.DecodeString(cryted)
	}
	shab := sha256.Sum256([]byte(key))
	k := pbkdf2.Key(shab[:], salt, iterationCount, 32, sha3.New256)
	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(err)
	}

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = pkcs7UnPadding(orig)

	return string(orig)
}

//补码
func pkcs7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
