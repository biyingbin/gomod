package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"testing"
	"time"
)

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(src, src)
	return src, nil

}

// 解密
func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}

func TestAES(t *testing.T) {

	var key = []byte("7a00d82790d4910a" + "7a326a55a48316f8")

	d := []byte("hello,aseadafasdfjalkjdfalkjklajkflajflk;jklfajfklajfkljs;akljfklsjfkaslf;askjfsl;fkjklfajkldfa")
	//key := []byte("hgfedcba87654321")
	fmt.Println("加密前:", string(d))
	x1, err := encryptAES(d, key)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("加密后:", string(x1))

	s := time.Now()

	for i := 0; i < 10000*10000; i++ {

		b := make([]byte, len(x1))
		copy(b, x1)

		x2, err := decryptAES(b, key)
		if err != nil {
			log.Fatalln(err)
		}
		_ = x2

	}

	fmt.Println(time.Now().Sub(s).Seconds())

	//fmt.Println("解密后:", string(x2))
}
