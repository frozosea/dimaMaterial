package crawler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang-test-task/meta/pkg/util"
	"time"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) JSONLogMiddleware(c *gin.Context) {
	start := time.Now()
	// Process Request
	c.Next()
	// Stop timer
	duration := util.GetDurationInMilliseconds(start)
	entry := log.WithFields(log.Fields{
		"duration": duration,
		"method":   c.Request.Method,
		"path":     c.Request.RequestURI,
		"status":   c.Writer.Status(),
		"referrer": c.Request.Referer(),
	})
	if c.Writer.Status() >= 500 {
		entry.Error(c.Errors.String())
	} else {
		entry.Info("")
	}

}
func (m *Middleware) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Cache-Control, Content-Type, DNT, If-Modified-Since, Keep-Alive, Origin, User-Agent, X-Requested-With, X-Real-Ip, Access-Control-Allow-Origin")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	if c.Request.Method == "OPTIONS" {
		c.Writer.Header().Set("Access-Control-Max-Age", "1728000")
		c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		c.Writer.Header().Set("Content-Length", "0")
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
