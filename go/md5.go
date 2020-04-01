/**
* @program: Go
*
* @description:md5不存在解密
*
* @author: Mr.chen
*
* @create: 2020-04-01 10:08
**/
package main
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func main() {
	params := map[string]interface{} {
		"name" : "mr.chen",
		"pwd"  : "123456",

	}
	fmt.Printf("sign: %s\n", createSign(params))
}

// MD5 方法
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

// 生成签名
func createSign(params map[string]interface{}) string {

	var key []string
	var str = ""
	for k := range params {
		key   = append(key, k)
	}
	sort.Strings(key)
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], params[key[i]])
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], params[key[i]])
		}
	}
	fmt.Println(str)
	// 自定义密钥
	var secret = "f7rbsEwiO5bl1aTd"

	// 自定义签名算法
	sign := MD5(MD5(str) + MD5(secret))
	return sign
}
