package requests

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

func createRequestBodyForMultipartFormWithFile(form map[string]string, fieldName, filename string, file []byte) (string, io.Reader, error) {
	buf := new(bytes.Buffer)
	bw := multipart.NewWriter(buf) // RBody writer

	for key, value := range form {
		p, err := bw.CreateFormField(key)
		if err != nil {
			return "", nil, err
		}
		if _, err := p.Write([]byte(value)); err != nil {
			return "", nil, err
		}
	}
	fw1, _ := bw.CreateFormFile(fieldName, filename)
	if _, err := io.Copy(fw1, bytes.NewBuffer(file)); err != nil {
		return "", nil, err
	}

	if err := bw.Close(); err != nil {
		return "", nil, err
	}
	return bw.FormDataContentType(), buf, nil
}

type Request struct {
	url           string
	method        string
	headers       map[string]string
	form          map[string]string
	query         map[string]string
	multipartForm *multipartForm
	body          []byte
	request       *http.Request
}

func New() *Request {
	return &Request{
		multipartForm: &multipartForm{},
	}
}

func (r *Request) Url(url string) IHttp {
	r.url = url
	return r
}
func (r *Request) Method(method string) IHttp {
	r.method = strings.ToUpper(method)
	return r
}
func (r *Request) Headers(headers map[string]string) IHttp {
	r.headers = headers
	return r
}
func (r *Request) MultipartForm(form map[string]string, fieldName, filename string, file []byte) IHttp {
	r.multipartForm = &multipartForm{
		Form:      form,
		FieldName: fieldName,
		Filename:  filename,
		File:      file,
	}
	return r
}
func (r *Request) Form(form map[string]string) IHttp {
	r.form = form
	return r
}
func (r *Request) Body(body []byte) IHttp {
	r.body = body
	return r
}
func (r *Request) Query(q map[string]string) IHttp {
	r.query = q
	return r
}
func (r *Request) addHeaders() {
	var mu sync.Mutex
	for key, value := range r.headers {
		mu.Lock()
		r.request.Header.Add(key, value)
		mu.Unlock()
	}
}
func (r *Request) getForm() url.Values {
	urlValues := make(url.Values)
	for key, value := range r.form {
		urlValues.Add(key, value)
	}
	return urlValues
}
func (r *Request) addForm() {
	r.request.PostForm = r.getForm()
}
func (r *Request) createEmtpyRequest(ctx context.Context) error {
	request, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(r.body))
	if err != nil {
		return err
	}
	request.WithContext(ctx)
	r.request = request
	return nil
}
func (r *Request) do(ctx context.Context) (*Response, error) {
	if err := r.createEmtpyRequest(ctx); err != nil {
		return nil, err
	}
	if len(r.form) != 0 {
		r.request.Form = r.getForm()
		r.request.PostForm = r.getForm()
	}
	if len(r.query) != 0 {
		for key, value := range r.query {
			q := r.request.URL.Query()
			q.Add(key, value)
			r.request.URL.RawQuery = q.Encode()
		}
	}
	if r.multipartForm.FieldName != "" {
		contentType, reader, err := createRequestBodyForMultipartFormWithFile(r.multipartForm.Form, r.multipartForm.FieldName, r.multipartForm.Filename, r.multipartForm.File)
		if err != nil {
			return nil, err
		}
		r.request, err = http.NewRequest(r.method, r.url, reader)
		if err != nil {
			return nil, err
		}
		r.request.Header.Add("Content-Type", contentType)
	}
	cli := &http.Client{}
	if len(r.headers) != 0 {
		r.addHeaders()
	}
	resp, err := cli.Do(r.request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	byteBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:          byteBody,
		Status:        resp.StatusCode,
		ContentType:   resp.Header.Get("Content-type"),
		ContentLength: resp.ContentLength,
	}, nil
}
func (r *Request) Do(ctx context.Context) (*Response, error) {
	if len(r.url) == 0 {
		return nil, &NoUrlError{}
	}
	if len(r.method) == 0 {
		return nil, &NoMethodError{}
	}
	return r.do(ctx)
}
