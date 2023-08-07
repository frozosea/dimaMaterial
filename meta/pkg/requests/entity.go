package requests

type Response struct {
	Body          []byte
	Status        int    `json:"status"`
	ContentType   string `json:"content-type"`
	ContentLength int64  `json:"content-length"`
}

type multipartForm struct {
	Form      map[string]string
	FieldName string
	Filename  string
	File      []byte
}
