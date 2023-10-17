package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin/blog-service/pkg/app"
	"github.com/gin/blog-service/pkg/errcode"
	"github.com/gin/blog-service/pkg/limiter"
)

func RateLimiter(l limiter.LimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
