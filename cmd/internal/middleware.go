package internal

import (
	"IpLimiter/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RateLimiterMiddleware struct {
	Limiter *pkg.Limiter
}

func NewRateLimiterMiddleware(limiter *pkg.Limiter) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		Limiter: limiter,
	}
}

func (m *RateLimiterMiddleware) handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		attemptCount, ttl, err := m.Limiter.Limit(ip)
		if err != nil {
			c.Abort()
			return
		}

		remaining := m.Limiter.Options.MaxAttempt - attemptCount
		if remaining < 0 {
			remaining = 0
		}
		c.Writer.Header().Set("X-RateLimit-Remaining", strconv.FormatInt(remaining, 10))
		c.Writer.Header().Set("X-RateLimit-Reset", strconv.Itoa(int(ttl.Seconds())))

		if remaining == 0 {
			c.AbortWithStatus(http.StatusTooManyRequests)
			fmt.Printf("Attempt too many times: %s \n", ip)
			return
		}

		m.Limiter.Hit(ip)
		c.Next()
	}
}
