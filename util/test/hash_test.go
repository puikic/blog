package test

import (
	"blog/util"
	"testing"
)

func TestMd5(t *testing.T) {
	util.Md5("test")
}
//go test -v ./util/test -run=^TestMd5$ -count=1