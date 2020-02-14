package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func buildParams(params map[string]string, request *http.Request) {
	q := request.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	request.URL.RawQuery = q.Encode()
}

func buildHeaders(headers map[string]string, request *http.Request) {
	for k, v := range headers {
		request.Header.Add(k, v)
	}
}

func Get(url string, params map[string]string, headers map[string]string) {
	client := &http.Client{}
	//提交请求
	request, err := http.NewRequest("GET", url, nil)
	buildParams(params, request)
	buildHeaders(headers, request)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(body))
}

func Post() {

}

func Put() {

}

func Option() {

}

func Head() {

}

func Patch() {

}

func Delete() {

}
