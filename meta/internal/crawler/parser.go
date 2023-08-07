package crawler

import (
	"golang.org/x/net/html"
	"io"
	"strings"
	"sync"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) GetTags(stringHTML string) ([]string, error) {
	reader := strings.NewReader(stringHTML)
	tokenizer := html.NewTokenizer(reader)
	return p.TraverseTags(tokenizer)
}
func (p *Parser) GetTagsCount(tags []string) map[string]int {
	dict := make(map[string]int)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, num := range tags {
		wg.Add(1)
		go func(tag string) {
			defer mu.Unlock()
			defer wg.Done()
			mu.Lock()
			dict[tag] = dict[tag] + 1
		}(num)
	}
	wg.Wait()
	return dict
}
func (p *Parser) TraverseTags(n *html.Tokenizer) ([]string, error) {
	var ar []string
	for {
		tt := n.Next()
		if tt == html.ErrorToken {
			if n.Err() == io.EOF {
				return ar, nil
			}
		}
		rawTag, _ := n.TagName()
		tag := string(rawTag)
		if tag != "" {
			ar = append(ar, tag)
		}
	}
}
