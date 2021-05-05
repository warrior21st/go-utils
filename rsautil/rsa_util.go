package rsautil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"

	"github.com/warrior21st/go-utils/base58util"
)

//加密-使用1024位n，明文最大117字节数据（pkcs1占11字节）
func RsaEncrypt(origData string, publicKey string) string {
	return rsaEncryptWithEncoding(origData, publicKey, false)
}

//解密-使用1024位n，明文最大117字节数据（pkcs1占11字节）
func RsaDecrypt(origData string, publicKey string) string {
	return rsaDecryptWithEncoding(origData, publicKey, false)
}

//加密-使用base58编码,使用1024位n，明文最大117字节数据（pkcs1占11字节）
func RsaEncryptBase58(origData string, publicKey string) string {
	return rsaEncryptWithEncoding(origData, publicKey, true)
}

//解密-使用base58编码,使用1024位n，明文最大117字节数据（pkcs1占11字节）
func RsaDecryptBase58(origData string, publicKey string) string {
	return rsaDecryptWithEncoding(origData, publicKey, true)
}

//加密-指定编码
func rsaEncryptWithEncoding(origData string, publicKey string, useBase58 bool) string {
	//解密pem格式的公钥
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		panic(errors.New("public key error"))
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	resBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	if err != nil {
		panic(err)
	}

	var result string
	if useBase58 {
		result = base58util.Base58Encode(resBytes)
	} else {
		result = base64.StdEncoding.EncodeToString(resBytes)
	}

	return result
}

//解密-指定编码
func rsaDecryptWithEncoding(cipherText string, privateKey string, useBase58 bool) string {
	//解密
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		panic(errors.New("private key error!"))
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// 解密
	var cipherBytes []byte
	if useBase58 {
		cipherBytes = base58util.Base58DecodeToBytes([]byte(cipherText))
	} else {
		cipherBytes, _ = base64.StdEncoding.DecodeString(cipherText)
	}
	resBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherBytes)
	if err != nil {
		panic(err)
	}

	return string(resBytes)
}

//生成1024位rsa密钥对，返回（私钥，公钥）
func GenRsaKey() (string, string) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	privateKeyStr := string(pem.EncodeToMemory(priBlock))

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	publicKeyStr := string(pem.EncodeToMemory(publicBlock))

	return privateKeyStr, publicKeyStr
}
