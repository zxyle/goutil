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

func Test_Strip(t *testing.T) {
	ret := str.Strip(" s ")
	if ret != "s" {
		panic("err")
	}
}

func Test_LStrip(t *testing.T) {
	ret := str.LStrip(" s ")
	if ret != "s " {
		panic("err")
	}
}

func Test_RStrip(t *testing.T) {
	ret := str.RStrip(" s ")
	if ret != " s" {
		panic("err")
	}
}

func Test_Join(t *testing.T) {
	ret := str.Join(",", []string{"hello", "world", "golang"})
	if ret != "hello,world,golang" {
		panic("err")
	}
}
