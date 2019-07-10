package http

import (
	"errors"
	"fmt"

	"github.com/valyala/fasthttp"
)

const (
	HttpMethodGet  = "GET"
	HttpMethodPost = "POST"
)

type HttpClient struct{}

func (*HttpClient) Get(url string, data []byte) ([]byte, error) {
	return http(HttpMethodGet, url, data)
}
func (*HttpClient) Post(url string, data []byte) ([]byte, error) {
	return http(HttpMethodPost, url, data)
}

func http(reqMethod, url string, repBody []byte) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	if reqMethod == HttpMethodGet {
		req.Header.SetMethod(HttpMethodGet)
	}
	if reqMethod == HttpMethodPost {
		req.Header.SetMethod(HttpMethodPost)
	}

	req.SetRequestURI(url)

	req.SetBody(repBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源

	if err := fasthttp.Do(req, resp); err != nil {
		return []byte{}, err
	}

	return resp.Body(), nil

}

func HttpGet(url string) ([]byte, error) {
	status, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		// 请求发生错误
		return []byte{}, err
	}

	if status != fasthttp.StatusOK {
		return []byte{}, errors.New(fmt.Sprintf("请求失败，status: %d", status))
	}

	return resp, nil
}

func HttpPost(url string, data map[string]string) ([]byte, error) {
	//填充表单，类似于net/url

	args := &fasthttp.Args{}
	for k, v := range data {
		args.Add(k, v)
	}

	status, resp, err := fasthttp.Post(nil, url, args)
	if err != nil {
		//请求发生错误
		return []byte{}, err
	}

	if status != fasthttp.StatusOK {
		return []byte{}, errors.New(fmt.Sprintf("请求失败，status: %d", status))
	}

	return resp, nil
}
