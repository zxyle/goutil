package main

import (
	"github.com/zxyle/goutils/str"
	"testing"
)

func Test_ZFill(t *testing.T) {
	a := str.PyStr("h")
	ret := a.ZFill(5)
	if ret != "0000h" {
		panic("err")
	}
}

func Test_Strip(t *testing.T) {
	a := str.PyStr(" s ")
	ret := a.Strip()
	if ret != "s" {
		panic("err")
	}
}

func Test_LStrip(t *testing.T) {
	a := str.PyStr(" s ")
	ret := a.LStrip()
	if ret != "s " {
		panic("err")
	}
}

func Test_RStrip(t *testing.T) {
	a := str.PyStr(" s ")
	ret := a.RStrip()
	if ret != " s" {
		panic("err")
	}
}

func Test_Join(t *testing.T) {
	a := str.PyStr(",")
	ret := a.Join([]string{"hello", "world", "golang"})
	if ret != "hello,world,golang" {
		panic("err")
	}
}
