/*
 * 纸喵软件
 * Copyright (c) 2017~2020 http://zhimiao.org All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2020/3/3 下午4:26
 */

package gutils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// MD5 md5 encryption
func MD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// ParseToken 解析jwtToken
func ParseToken(tokenString string, secret []byte) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
				return "", fmt.Errorf("过期了")
			}
			if claims.Issuer != "zhimiao" {
				return "", fmt.Errorf("非法来源的签名")
			}
			return claims.Subject, nil
		}
	}
	return "", err
}

// CreateToken 生成jwtToken
func CreateToken(subject string, expire time.Duration, secret []byte) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   subject,
		ExpiresAt: time.Now().Add(expire).Unix(),
		Issuer:    "zhimiao",
	})
	token, err := tokenClaims.SignedString(secret)
	return token, err
}

// PasswordHash 密码生成
func PasswordHash(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		logrus.Warn(err.Error())
	}
	return string(hash)
}

// PasswordVerify 密码验证
func PasswordVerify(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err != nil {
		logrus.Warn(err.Error())
		return false
	}
	return true
}

// ---------------DES加密  解密--------------------

// EncyptogAES AES加密
func EncyptogAES(src, key string) string {
	s := []byte(src)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		print(err.Error())
		return ""
	}
	blockSize := block.BlockSize()
	paddingCount := blockSize - len(s)%blockSize
	// 填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newStr := append(s, paddingStr...)
	blockMode := cipher.NewCBCEncrypter(block, []byte(key))
	blockMode.CryptBlocks(newStr, newStr)
	return string(newStr)

}

// DecrptogAES AES解密
func DecrptogAES(src, key string) string {
	s := []byte(src)
	k := []byte(key)
	block, err := aes.NewCipher(k)
	if err != nil {
		return ""
	}
	blockMode := cipher.NewCBCDecrypter(block, k)
	blockMode.CryptBlocks(s, s)
	n := len(s)
	count := int(s[n-1])
	return string(s[:n-count])
}

// ---------------RSA 密钥对生成 加密 解密--------------------

// GenerateRSAKey 生成RSA私钥和公钥，保存到文件中
func GenerateRSAKey(bits int) (pubKey, priKey []byte, err error) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	priKey = pem.EncodeToMemory(&privateBlock)

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return
	}
	//pem格式编码
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pubKey = pem.EncodeToMemory(&publicBlock)
	return
}

// RSAEncrypt RSA加密
func RSAEncrypt(plainText []byte, key []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode(key)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	//返回密文
	return cipherText, nil
}

// RSADecrypt RSA解密
func RSADecrypt(cipherText []byte, key []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode(key)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText, nil
}
