package main

import (
	"github.com/zxyle/goutil/str"
	"testing"
)

func Test_Zfill(t *testing.T) {
	ret := str.Zfill("h", 5)
	if ret != "0000h" {
		panic("err")
	}
}
