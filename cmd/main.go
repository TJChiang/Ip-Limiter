package main

import (
	"IpLimiter/pkg"
	"fmt"
)

func main() {
	ip := "127.0.0.1"
	limiter, err := pkg.NewLimiter("ip_block:")
	if err != nil {
		panic(err)
	}

	hitCount, err := limiter.Hit(ip)
	if err != nil {
		panic(err)
	}

	attemptCount, ttl, err := limiter.Limit(ip)

	fmt.Println(hitCount, attemptCount, ttl)
}
