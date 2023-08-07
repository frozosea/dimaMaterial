package crawler

import (
	"context"
	"github.com/go-playground/assert/v2"
	http "golang-test-task/meta/pkg/requests"
	"os"
	"testing"
)

type LoggerMockup struct {
}

func NewLoggerMockup() *LoggerMockup {
	return &LoggerMockup{}
}

func (l LoggerMockup) InfoLog(_ string) {
	return
}

func (l LoggerMockup) ExceptionLog(_ string) {
	return
}

func (l LoggerMockup) WarningLog(_ string) {
	return
}

func (l LoggerMockup) FatalLog(_ string) {
	return
}

func TestWithHttpMoch(t *testing.T) {
	if !testing.Short() {
		t.Skip()
	}

	p := NewParser()
	serviceLogger := NewLoggerMockup()
	requestSender := http.NewRequestMockUp(200, func(r http.RequestMockUp) ([]byte, error) {
		return os.ReadFile("test_data/example.txt")
	})
	service := NewService(requestSender, p, serviceLogger)
	ctx := context.Background()

	r, err := service.Do(ctx, []string{"https://example.com/", "https://example.com/", "https://example.com/"})
	assert.Equal(t, err, nil)
	assert.Equal(t, len(r), 3)
	for _, v := range r {
		assert.Equal(t, v.Url, "https://example.com/")
		assert.Equal(t, len(v.Elements), 11)
	}
}
