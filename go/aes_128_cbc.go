/**
* @program: Go
*
* @description:加密前： {"bindKey":"8IIYer87yte0vi7p"}加密后： SnnV3T1CwGb8n0qFGOlYikIHcBn/OmUVf6uHgYJS4Ks=解密后： {"bindKey":"8IIYer87yte0vi7p"}
*
* @author: Mr.chen
*
* @create: 2020-04-01 09:54
**/

package main
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
)
type aqara struct {
	BindKey string `json:"bindKey"`  // BindKey首字母必须为大写
}
// 定义密钥和 iv 向量
const (
	key = "f7rbsEwiO5bl1aTd"
	iv  = "f7rbsEwiO5bl1aTd" //Go 的 iv 向量为必传参数
)

func main() {
	a := &aqara{}
	a.BindKey = "8IIYer87yte0vi7p"
	data,_:= json.Marshal(a)

	str := string(data)
	fmt.Println("加密前：", str)

	encrypted, err := encrypt(str, []byte(key))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("加密后：", encrypted)
	}

	decrypted, err := decrypt(encrypted, []byte(key))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("解密后：", decrypted)
	}
}

// 加密
func encrypt (encryptStr string, key []byte) (string, error) {
	IV := []byte(iv);
	if IV != nil && len(IV) != 16 {
		return "", errors.New("iv偏移量必须为16位数")
	}
	encryptBytes := []byte(encryptStr)
	block, err   := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encryptBytes = PKCS5Padding(encryptBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// 解密
func decrypt (decryptStr string, key []byte) (string, error) {
	decryptBytes, err := base64.StdEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	decrypted := make([]byte, len(decryptBytes))

	blockMode.CryptBlocks(decrypted, decryptBytes)
	decrypted = PKCS5UnPadding(decrypted)
	return string(decrypted), nil
}

func PKCS5Padding (cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5UnPadding (decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}
