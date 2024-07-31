package util

import (
	"math/rand"
)

func isUpperChar(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func UpperLowerExchange(b byte) byte {
	return b ^ ' '
}

func Camel2Snake(s string) string {
	//testName->test_name
	//Test->test
	//
	if len(s) == 0 {
		return ""
	}
	t := make([]byte, 0, len(s)+4)
	if isUpperChar(s[0]) {
		t = append(t, UpperLowerExchange(s[0]))
	} else {
		t = append(t, s[0])
	}
	for i := 1; i < len(s); i++ {
		if isUpperChar(s[i]) {
			t = append(t, '_', UpperLowerExchange(s[i]))
		} else {
			t = append(t, s[i])
		}
	}
	return string(t)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ测试中文")

// 生成随机字符串
func RandStringRunes(n int) string {
	r := make([]rune, n) // len(b)==cap(b)==n
	//fmt.Println(r)
	l := len(letterRunes)
	for i := range r {
		r[i] = letterRunes[rand.Intn(l)]
	}
	return string(r)
}
