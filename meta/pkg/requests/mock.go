package requests

import (
	"context"
	"strings"
)

type RequestMockUp struct {
	WrapFunc   func(r RequestMockUp) ([]byte, error)
	StatusCode int
	RUrl       string
	RMethod    string
	RHeaders   map[string]string
	RForm      map[string]string
	RQuery     map[string]string
	RMultipart *multipartForm
	RBody      []byte
}

func NewRequestMockUp(statusCode int, wrapFunc func(r RequestMockUp) ([]byte, error)) *RequestMockUp {
	return &RequestMockUp{WrapFunc: wrapFunc, StatusCode: statusCode}
}
func (r *RequestMockUp) Url(url string) IHttp {
	r.RUrl = url
	return r
}
func (r *RequestMockUp) Method(method string) IHttp {
	r.RMethod = strings.ToUpper(method)
	return r
}
func (r *RequestMockUp) Headers(headers map[string]string) IHttp {
	r.RHeaders = headers
	return r
}
func (r *RequestMockUp) MultipartForm(form map[string]string, fieldName, filename string, file []byte) IHttp {
	r.RMultipart = &multipartForm{
		Form:      form,
		FieldName: fieldName,
		Filename:  filename,
		File:      file,
	}
	return r
}
func (r *RequestMockUp) Form(form map[string]string) IHttp {
	r.RForm = form
	return r
}
func (r *RequestMockUp) Body(body []byte) IHttp {
	r.RBody = body
	return r
}
func (r *RequestMockUp) Query(q map[string]string) IHttp {
	r.RQuery = q
	return r
}

func (r RequestMockUp) Do(_ context.Context) (*Response, error) {
	body, err := r.WrapFunc(r)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:          body,
		Status:        r.StatusCode,
		ContentType:   "",
		ContentLength: 0,
	}, nil
}

type UserAgentGeneratorMockUp struct {
}

func NewUserAgentGeneratorMockUp() *UserAgentGeneratorMockUp {
	return &UserAgentGeneratorMockUp{}
}

func (u *UserAgentGeneratorMockUp) Generate() string {
	return ""
}
