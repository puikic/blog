package util

import (
	"fmt"
	"math/rand"
)

func isUpperChar(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

// func Camel2Snake(s string) string {
// 	//testName->test_name
// 	//Test->test
// 	//
// 	if len(s) == 0 {
// 		return ""
// 	}
// 	t := make([]string, 0)
// 	if isUpperChar(s[0]) {

// 	}
// 	return ""
// }

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ测试中文")

// 生成随机字符串
func RandStringRunes(n int) string {
	r := make([]rune, n) // len(b)==cap(b)==n
	fmt.Println(r)
	l := len(letterRunes)
	for i := range r {
		r[i] = letterRunes[rand.Intn(l)]
	}
	return string(r)
}
