package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

func RateLimit() func(c *gin.Context) {
	bucket := ratelimit.NewBucket(time.Second, 1700)
	return func(c *gin.Context) {
		// 如果取不到令牌就返回响应
		if bucket.TakeAvailable(1) < 1 {
			c.Abort()
			return
		}
		// 取到令牌放行
		c.Next()
	}
}
