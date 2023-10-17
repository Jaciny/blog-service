package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

func NewMethodLimiter() LimiterInterface {
	return MethodLimiter{
		&Limiter{LimiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

type MethodLimiter struct {
	*Limiter
}

func (l MethodLimiter) Key(c *gin.Context) string {
	//TODO implement me
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	//TODO implement me
	bucket, ok := l.LimiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
	//TODO implement me
	for _, rule := range rules {
		if _, ok := l.LimiterBuckets[rule.Key]; !ok {
			l.LimiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}

	return l
}
