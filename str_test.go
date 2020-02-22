package main

import (
	"github.com/zxyle/goutil/string"
	"testing"
)

func Test_Zfill(t *testing.T) {
	ret := string.Zfill("h", 5)
	if ret != "0000h" {
		panic("err")
	}
}
