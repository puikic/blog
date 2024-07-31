package test

import (
	"blog/util"
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	str := util.RandStringRunes(30)
	fmt.Println(str)
}

func TestCamel2Snake(t *testing.T) {
	str := util.Camel2Snake("Abc")
	if str != "abc" {
		fmt.Println(str)
		t.Fail()
	}
	str = util.Camel2Snake("ABC")
	if str != "a_b_c" {
		fmt.Println(str)
		t.Fail()
	}
	str = util.Camel2Snake("AbcEfg")
	if str != "abc_efg" {
		fmt.Println(str)
		t.Fail()
	}
	str = util.Camel2Snake("abcEF")
	if str != "abc_e_f" {
		fmt.Println(str)
		t.Fail()
	}
	str = util.Camel2Snake("abcE")
	if str != "abc_e" {
		fmt.Println(str)
		t.Fail()
	}
}

// go test -v ./util/test -run=^TestRandString$ -count=1
// go test -v ./util/test -run=^TestCamel2Snake$ -count=1
