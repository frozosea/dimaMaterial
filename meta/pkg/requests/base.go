package requests

import "context"

type IHttp interface {
	Url(url string) IHttp
	Method(method string) IHttp
	Headers(headers map[string]string) IHttp
	Form(form map[string]string) IHttp
	MultipartForm(form map[string]string, fieldName, filename string, file []byte) IHttp
	Body(body []byte) IHttp
	Query(q map[string]string) IHttp
	Do(ctx context.Context) (*Response, error)
}
