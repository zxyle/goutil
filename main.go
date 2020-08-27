package main

import "github.com/zxyle/goutils/request"

func main() {
	request.Get("http://httpbin.org/ip", map[string]string{}, map[string]string{})
}
