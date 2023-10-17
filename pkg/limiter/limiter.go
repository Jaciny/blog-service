package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

//接口限流

type LimiterInterface interface {
	//获取对应的限流器的键值对名称。
	Key(c *gin.Context) string

	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterInterface
}

type Limiter struct {
	LimiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration //间隔多久时间放 N 个令牌
	Capacity     int64         //令牌桶的容量
	Quantum      int64         //每次到达间隔时间后所放的具体令牌数量 N
}
