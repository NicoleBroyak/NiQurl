package redishandler

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestIsValidSettingTrue(t *testing.T) {
	boolean := isValidSetting("SHORT_URL_LEN")
	if !boolean {
		t.Fatalf(`Error: %v want match for true`, boolean)
	}
}

func TestIsValidSettingFalse(t *testing.T) {
	urllen := GetSetting("SHORT_URL_LEN")
	client.Del(context, "SHORT_URL_LEN")
	boolean := isValidSetting("SHORT_URL_LEN")
	ChangeSetting("SHORT_URL_LEN", urllen)
	if boolean {
		t.Fatalf(`Error: %v want match for false`, boolean)
	}
}

func TestExistsValueInSortedSetTrue(t *testing.T) {
	client.ZAdd(context, "testset", &redis.Z{
		Score:  1,
		Member: "testvalue1",
	})
	client.ZAdd(context, "testset", &redis.Z{
		Score:  2,
		Member: "testvalue2",
	})
	boolean := existsValueInSortedSet("testset", "testvalue2")
	if !boolean {
		t.Fatalf(`Error: %v want match for true`, boolean)
	}
	client.Del(context, "testset")
}

func TestExistsValueInSortedSetFalse(t *testing.T) {
	client.ZAdd(context, "testset", &redis.Z{
		Score:  1,
		Member: "testvalue1",
	})
	boolean := existsValueInSortedSet("testset", "testvalue2")
	if boolean {
		t.Fatalf(`Error: %v want match for false`, boolean)
	}
	client.Del(context, "testset")
}

func TestIsUserOnWaitTimeTrue(t *testing.T) {
	client.Set(context, "testuser", true, time.Duration(30000000))
	client.Set(context, "anothertestuser", true, time.Duration(30000000))
	boolean := IsUserOnWaitTime("anothertestuser")
	if !boolean {
		t.Fatalf(`Error: %v want match for false`, boolean)
	}
	client.Del(context, "testuser", "anothertestuser")
}

func TestIsUserOnWaitTimeFalse(t *testing.T) {
	client.Set(context, "testuser", true, time.Duration(30000000))
	boolean := IsUserOnWaitTime("anothertestuser")
	if boolean {
		t.Fatalf(`Error: %v want match for false`, boolean)
	}
	client.Del(context, "testuser")
}

func TestExistsUser(t *testing.T) {
	userCount := client.ZCount(context, "username", "-inf", "inf")
	client.ZAdd(context, "username", &redis.Z{
		Score:  float64(userCount.Val()) + 1.0,
		Member: "RandomTestUsernameNiqURL",
	})
	boolean := existsUser("RandomTestUsernameNiqURL")
	client.ZPopMax(context, "username", 0)
	if !boolean {
		t.Fatalf(`Error: %v want match for true`, boolean)
	}
	client.ZPopMax(context, "username", 0)
}

func TestExistsLongURL(t *testing.T) {
	urlCount := client.ZCount(context, "longurl", "-inf", "inf")
	client.ZAdd(context, "longurl", &redis.Z{
		Score:  float64(urlCount.Val()) + 1.0,
		Member: "RandomNiqURLwebsite.com",
	})
	boolean := ExistsLongURL("RandomNiqURLwebsite.com")
	if !boolean {
		t.Fatalf(`Error: %v want match for true`, boolean)
	}
	client.ZPopMax(context, "shorturl", 0)
}

func TestExistsShortURL(t *testing.T) {
	urlCount := client.ZCount(context, "shorturl", "-inf", "inf")
	client.ZAdd(context, "shorturl", &redis.Z{
		Score:  float64(urlCount.Val()) + 1.0,
		Member: "L29x5a",
	})
	boolean := ExistsShortURL("L29x5a")
	if !boolean {
		t.Fatalf(`Error: %v want match for true`, boolean)
	} else {
		printExistingShortURL("L29x5a")
	}
	client.ZPopMax(context, "shorturl", 0)
}
