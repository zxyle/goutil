package main

import "goutil/request"

func main() {
	request.Get("http://httpbin.org/ip", map[string]string{}, map[string]string{})
}
