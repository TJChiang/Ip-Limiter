package tests

import (
	"IpLimiter/pkg"
	"testing"
)

var testingIpAddr string = "255.255.255.255"

func TestLimiterHit(t *testing.T) {
	testingLimiter, _ := pkg.NewLimiter("testing")
	testingLimiter.Hit(testingIpAddr)
	value, err := testingLimiter.Rdb.Get(testingLimiter.Ctx, "testing:"+testingIpAddr).Result()
	if err != nil {
		t.Error(err)
		testingLimiter.Rdb.Del(testingLimiter.Ctx, "testing:"+testingIpAddr)
		return
	}

	if value != "1" {
		t.Error("Invalid value.")
	}

	testingLimiter.Rdb.Del(testingLimiter.Ctx, "testing:"+testingIpAddr)
}
