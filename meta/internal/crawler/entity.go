package crawler

type Meta struct {
	Status        int    `json:"status"`
	ContentType   string `json:"content-type"`
	ContentLength int64  `json:"content-length"`
}
type Element struct {
	TagName string `json:"tag-name"`
	Count   int    `json:"count"`
}
type Response struct {
	Url      string     `json:"url"`
	Meta     *Meta      `json:"meta"`
	Elements []*Element `json:"elements"`
}
