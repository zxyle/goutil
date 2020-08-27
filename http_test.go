package main

import (
	"fmt"
	"github.com/zxyle/goutils/request"
	"testing"
)

func Test_Get(t *testing.T) {
	resp := request.Get("http://httpbin.org/ip", map[string]string{}, map[string]string{})
	fmt.Println(resp.Reason)
}
