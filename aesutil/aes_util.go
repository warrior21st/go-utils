package aesutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"

	"github.com/warrior21st/go-utils/base58util"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/sha3"
)

const iterationCount int = 2009

//aes加密
func AesEncrypt(orig string, key string, salt string, useBase58 bool) string {
	// 转成字节数组
	origData := []byte(orig)
	shaBytes := sha256.Sum256([]byte(key))
	saltBytes := sha256.Sum256([]byte(salt))
	k := pbkdf2.Key(shaBytes[:], saltBytes[:], iterationCount, 32, sha3.New256)
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
func AesDecrypt(cryted string, key string, salt string, useBase58 bool) string {
	var crytedByte []byte
	// 转成字节数组
	if useBase58 {
		crytedByte = base58util.Base58DecodeToBytes([]byte(cryted))
	} else {
		crytedByte, _ = base64.StdEncoding.DecodeString(cryted)
	}
	shaBytes := sha256.Sum256([]byte(key))
	saltBytes := sha256.Sum256([]byte(salt))
	k := pbkdf2.Key(shaBytes[:], saltBytes[:], iterationCount, 32, sha3.New256)
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

func AesEncryptWithIterations(orig string, key string, salt string, iterations int) string {
	// 转成字节数组
	origData := []byte(orig)
	shaBytes := sha256.Sum256([]byte(key))
	saltBytes := sha256.Sum256([]byte(salt))
	k := pbkdf2.Key(shaBytes[:], saltBytes[:], iterations, 32, sha3.New256)
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

	return base64.StdEncoding.EncodeToString(cryted)
}

//aes解密
func AesDecryptWithIterations(cryted string, key string, salt string) string {
	var crytedByte []byte
	// 转成字节数组
	crytedByte, _ = base64.StdEncoding.DecodeString(cryted)
	shaBytes := sha256.Sum256([]byte(key))
	saltBytes := sha256.Sum256([]byte(salt))
	k := pbkdf2.Key(shaBytes[:], saltBytes[:], iterationCount, 32, sha3.New256)
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

//使用AES-GCM加密(nonce必须为12位)
func AesGCMEncrypt(plaintext, key, nonce []byte) []byte {
	shaBytes := sha256.Sum256(key)
	block, err := aes.NewCipher(shaBytes[:])
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext
}

//使用AES-GCM解密(nonce必须为12位)
func AesGCMDecrypt(ciphertext, key, nonce []byte) []byte {
	shaBytes := sha256.Sum256(key)
	block, err := aes.NewCipher(shaBytes[:])
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}
