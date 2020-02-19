package request

import (
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

type Response struct {
	// 'apparent_encoding', 'close', 'connection', , 'cookies', 'elapsed', 'encoding', 'headers', 'history', 'is_permanent_redirect', 'is_redirect', 'iter_content', 'iter_lines', 'json', 'links', 'next',  'raise_for_status', 'raw', , 'request']
	Content    []byte
	Reason     string
	Ok         bool
	StatusCode int
	Url        string
	Text       string
	Encoding   string
}

func Get(url string, params map[string]string, headers map[string]string) Response {
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
	resp := Response{
		Content:    body,
		Reason:     "OK",
		Ok:         true,
		StatusCode: response.StatusCode,
		Url:        url,
		Text:       string(body),
		Encoding:   "",
	}
	return resp
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
