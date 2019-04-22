package erroremail

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-ops/modules/email"
	"net/http"
	"net/http/httputil"
	"time"
)

func timeFormat(t time.Time) string {
	var timeString = t.Format("2006/01/02 - 15:04:05")
	return timeString
}

func ErrEmailWriter() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				errMsg := fmt.Sprintf("[Recovery] %s panic recovered:\n%s\n%s\n%s%s", timeFormat(time.Now()), string(httprequest), err)
				email.SendEmail(nil, "request error", errMsg, "text")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}