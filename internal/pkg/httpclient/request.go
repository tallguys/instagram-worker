package httpclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Request http request
type Request struct {
	Scheme  string
	Host    string
	Path    string
	Query   url.Values
	Header  http.Header
	Body    []byte
	Cookies []*http.Cookie
}

// Response http Response
type Response struct {
	Body       string
	Header     http.Header
	StatusCode int
	Cookies    []*http.Cookie
}

// RequestError request error
type RequestError struct {
	URL string
	Msg string
}

func (e RequestError) Error() string {
	return fmt.Sprintf("failed to request %s with error %s", e.URL, e.Msg)
}

func setHeader(header http.Header, req *http.Request) {
	for key, value := range header {
		req.Header.Add(key, value[0])
	}
}

func setCookies(cookies []*http.Cookie, req *http.Request) {
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
}

// Get http get method
func (req Request) Get() (Response, error) {
	client := http.Client{}
	urlStr := (&url.URL{
		Scheme:   req.Scheme,
		Host:     req.Host,
		Path:     req.Path,
		RawQuery: req.Query.Encode(),
	}).String()

	request, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	setHeader(req.Header, request)
	setCookies(req.Cookies, request)

	res, err := client.Do(request)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	bodyStr := string(bodyBytes)

	return Response{
		Header:     res.Header,
		Body:       bodyStr,
		StatusCode: res.StatusCode,
		Cookies:    res.Cookies(),
	}, nil
}

// Delete http delete method
func (req Request) Delete() (Response, error) {
	client := http.Client{}
	urlStr := (&url.URL{
		Scheme:   req.Scheme,
		Host:     req.Host,
		Path:     req.Path,
		RawQuery: req.Query.Encode(),
	}).String()

	request, err := http.NewRequest(http.MethodDelete, urlStr, nil)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	setHeader(req.Header, request)
	setCookies(req.Cookies, request)

	res, err := client.Do(request)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	bodyStr := string(bodyBytes)

	return Response{
		Header:     res.Header,
		Body:       bodyStr,
		StatusCode: res.StatusCode,
		Cookies:    res.Cookies(),
	}, nil
}

// Post http post method
func (req Request) Post() (Response, error) {
	client := http.Client{}
	urlStr := (&url.URL{
		Scheme:   req.Scheme,
		Host:     req.Host,
		Path:     req.Path,
		RawQuery: req.Query.Encode(),
	}).String()

	request, err := http.NewRequest(http.MethodPost, urlStr, bytes.NewBuffer(req.Body))
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	setHeader(req.Header, request)
	setCookies(req.Cookies, request)

	res, err := client.Do(request)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Response{}, RequestError{URL: urlStr, Msg: err.Error()}
	}
	bodyStr := string(bodyBytes)

	return Response{
		Header:     res.Header,
		Body:       bodyStr,
		StatusCode: res.StatusCode,
		Cookies:    res.Cookies(),
	}, nil
}
