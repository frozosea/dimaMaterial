package crawler

import (
	"context"
	"fmt"
	"golang-test-task/meta/pkg/logging"
	"golang-test-task/meta/pkg/requests"
	"sync"
)

type Service struct {
	http   requests.IHttp
	parser *Parser
	logger logging.ILogger
}

func NewService(http requests.IHttp, parser *Parser, logger logging.ILogger) *Service {
	return &Service{http: http, parser: parser, logger: logger}
}

func (s *Service) getElements(stringHTML string) ([]*Element, error) {
	var elements []*Element
	allTags, err := s.parser.GetTags(stringHTML)
	if err != nil {
		return nil, err
	}
	countTags := s.parser.GetTagsCount(allTags)
	var wg sync.WaitGroup
	var mu sync.Mutex
	for tag, count := range countTags {
		wg.Add(1)
		go func(tag string, count int) {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			elements = append(elements, &Element{
				TagName: tag,
				Count:   count,
			})
		}(tag, count)

	}
	wg.Wait()
	return elements, nil
}
func (s *Service) walkUrlAndGetResponse(ctx context.Context, url string) (*Response, error) {
	response, err := s.http.Url(url).Method("GET").Do(ctx)
	if err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`send request to url: %s error is: %s`, url, err.Error()))
		return nil, err
	}
	tags, err := s.getElements(string(response.Body))
	if err != nil {
		go s.logger.ExceptionLog(fmt.Sprintf(`get tags for url: %s error is %s`, url, err.Error()))
		return nil, err
	}
	return &Response{
		Url: url,
		Meta: &Meta{
			Status:        response.Status,
			ContentType:   response.ContentType,
			ContentLength: response.ContentLength,
		},
		Elements: tags,
	}, nil
}
func (s *Service) Do(ctx context.Context, urls []string) ([]*Response, error) {
	var responses []*Response
	ctxWithCancel, cancel := context.WithCancel(ctx)
	errChan := make(chan error, len(urls))
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, url := range urls {
		wg.Add(1)
		go func(ctx context.Context, url string, c context.CancelFunc) {
			defer wg.Done()
			defer mu.Unlock()
			mu.Lock()
			r, err := s.walkUrlAndGetResponse(ctx, url)
			if err != nil {
				c()
				errChan <- err
				return
			}
			responses = append(responses, r)
		}(ctxWithCancel, url, cancel)
		wg.Wait()
	}
	select {
	case <-ctx.Done():
		cancel()
		return responses, nil
	case err := <-errChan:
		cancel()
		close(errChan)
		return responses, err
	default:
		cancel()
		return responses, nil
	}
}
