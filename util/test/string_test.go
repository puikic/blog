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
// go test -v ./util/test -run=^TestRandString$ -count=1