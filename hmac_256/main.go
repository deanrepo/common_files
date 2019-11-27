package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// //字符串排序
	// names := []string{"Hello", "World", "private", "folders", "Users", "workspace"}
	// for _, value := range names {
	// 	fmt.Println(value)
	// }
	// fmt.Println("============================")
	// sort.Strings(names)
	// for _, value := range names {
	// 	fmt.Println(value)
	// }
	// fmt.Println("============================")

	values := make([]string, 0)
	values = append(values, "1AF5FDCE43365A5E1BEE3038A289B5FA", "10001", strconv.Itoa(1567674348))
	for _, value := range values {
		fmt.Println(value)
	}
	fmt.Println("============================")
	sort.Strings(values)
	for _, value := range values {
		fmt.Println(value)
	}
	fmt.Println("============================")
	ret := strings.Join(values, "")
	fmt.Println(ret)
	cmpSig := "1e9d9cd76643302b8025116c0b700763e1c3f171869884499391c8445df4b54a"
	appSecret := "aCd0LLujG09M91zC"
	conRet := ComputeHmacSha256(ret, appSecret)
	fmt.Println("===>conRet: ", conRet)
	fmt.Println(cmpSig == conRet)

}

// ValidMAC reports whether messageMAC is a valid HMAC tag for message.
func ValidMAC(message, messageMAC, key []byte) bool {
	fmt.Println("===>", string(messageMAC))
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	fmt.Println(string(expectedMAC))
	return hmac.Equal(messageMAC, expectedMAC)
}

func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	//	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	return sha
	// return base64.StdEncoding.EncodeToString([]byte(sha))
}
