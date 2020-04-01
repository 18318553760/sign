/**
* @program: Go
*
* @description:rsa加密解密算法，支付宝支付常用，运行时候请把idea的配置output directory和working directory目录到当前的目录下，否则无法找到文件，或者直接在当前的目录下运行go build
*
* @author: Mr.chen
*
* @create: 2020-04-01 10:12
**/
package main
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
)

// 公钥加密
func RsaEncrypt(encryptStr string, path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 读取文件内容
	info, _ := file.Stat()
	buf := make([]byte,info.Size())
	file.Read(buf)

	// pem 解码
	block, _ := pem.Decode(buf)

	// x509 解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// 类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//对明文进行加密
	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(encryptStr))
	if err != nil {
		return "", err
	}

	//返回密文
	return base64.StdEncoding.EncodeToString(encryptedStr), nil
}

// 私钥解密
func RsaDecrypt(decryptStr string, path string) (string, error) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 获取文件内容
	info, _ := file.Stat()
	buf := make([]byte,info.Size())
	file.Read(buf)

	// pem 解码
	block, _ := pem.Decode(buf)

	// X509 解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	decryptBytes, err := base64.StdEncoding.DecodeString(decryptStr)

	//对密文进行解密
	decrypted, _ := rsa.DecryptPKCS1v15(rand.Reader,privateKey,decryptBytes)

	//返回明文
	return string(decrypted), nil
}
type aqara struct {
	BindKey string `json:"bindKey"`  // BindKey首字母必须为大写
}
func main() {
	a := &aqara{}
	a.BindKey = "8IIYer87yte0vi7p"
	data,_:= json.Marshal(a)

	str := string(data)
	fmt.Println("加密前：", str)

	encrypted, err := RsaEncrypt(str,"public.pem")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("加密后：", encrypted)
	}

	decrypted, err := RsaDecrypt(encrypted, "private.pem")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("解密后：", decrypted)
	}
}
