package redishandler

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

func TestGetValueFromSortedSetCorrect(t *testing.T) {
	sortedSets := []string{"testurlset", "shorturl", "longurl"}
	for _, sortedSet := range sortedSets {
		client.ZAdd(context, sortedSet, &redis.Z{
			Score:  10000,
			Member: "firsturl",
		},
			&redis.Z{
				Score:  10001,
				Member: "secondurl",
			},
			&redis.Z{
				Score:  10002,
				Member: "thirdurl",
			})
	}
	url := getValueFromSortedSet("testurlset", -2)
	if url != "secondurl" {
		t.Fatalf(`Error: %q want match for "secondurl"`, url)
	}
	shorturl := getShortURL(-3)
	if shorturl != "firsturl" {
		t.Fatalf(`Error: %q want match for "secondurl"`, shorturl)
	}
	longurl := GetLongURL(-1)
	if longurl != "thirdurl" {
		t.Fatalf(`Error: %q want match for "thirdurl"`, longurl)
	}
	client.Del(context, "testurlset")
	client.ZPopMax(context, "shorturl", 3)
	client.ZPopMax(context, "longurl", 3)
}

func TestGetValueFromSortedSetInvalid(t *testing.T) {
	client.ZAdd(context, "testurlset", &redis.Z{
		Score:  0,
		Member: "firsturl",
	},
		&redis.Z{
			Score:  1,
			Member: "secondurl",
		},
		&redis.Z{
			Score:  2,
			Member: "thirdurl",
		})
	url := getValueFromSortedSet("testurllist", 4)
	if url != "" {
		t.Fatalf(`Error: %q want match for empty string`, url)
	}
	client.Del(context, "testurlset")
}

func TestGetValueFromListInvalid(t *testing.T) {
	lenOfList := client.LLen(context, "createdby").Val()
	author := getValueFromList("createdby", lenOfList-1)
	if author != "" {
		t.Fatalf(`Error: %q want match for emptystring`, author)
	}
}

func TestGetValueFromListCorrect(t *testing.T) {
	client.RPush(context, "createdby", "Username23")
	lenOfList := client.LLen(context, "createdby").Val()
	author := GetURLAuthor(lenOfList - 1)
	client.RPop(context, "createdby")
	if author != "Username23" {
		t.Fatalf(`Error: %q want match for "Username23"`, author)
	}
}

func TestGetRandomUser(t *testing.T) {
	client.Get(context, "USER_COUNT").Int64()
	username := GetRandomUser()
	query, _ := client.ZScan(context, "username", 0, username, 0).Val()
	if len(query) == 0 {
		t.Fatalf(`Query error`)
	}
	if query[0] != username {
		t.Fatalf(`Error: %q want match for %q`, username, query[0])
	}
}

func TestGetSetting(t *testing.T) {
	urlCountFromGet, err := client.Get(context, "URL_COUNT").Int()
	if err != nil {
		t.Fatalf("Error from redis client")
	}
	urlCountGetSetting := GetSetting("URL_COUNT")
	if urlCountGetSetting != urlCountFromGet {
		t.Fatalf(`Error: %v want match for %v`, urlCountFromGet, urlCountGetSetting)
	}
}

func TestGetIndexOfValueFromSortedSet(t *testing.T) {
	client.ZAdd(context, "testurlset", &redis.Z{
		Score:  0,
		Member: "firsturl",
	},
		&redis.Z{
			Score:  1,
			Member: "secondurl",
		},
		&redis.Z{
			Score:  2,
			Member: "thirdurl",
		})
	index, _ := getIndexOfValueFromSortedSet("testurlset", "thirdurl")
	if index != 2 {
		t.Fatalf(`Error: %q want match for %q`, index, 2)
	}
	client.Del(context, "testurlset")
}

func TestGetIndexOfShortURL(t *testing.T) {
	urlCount := float64(client.ZCount(context, "shorturl", "-inf", "inf").Val())
	client.ZAdd(context, "shorturl", &redis.Z{
		Score:  urlCount,
		Member: "2a0xaz2",
	},
	)
	index, _ := GetIndexOfShortURL("2a0xaz2")
	if index != int64(urlCount) {
		t.Fatalf(`Error: %v want match for %v`, index, urlCount)
	}
	client.ZPopMax(context, "shorturl", 0)
}

func TestGetIndexOfLongURL(t *testing.T) {
	urlCount := float64(client.ZCount(context, "longurl", "-inf", "inf").Val())
	client.ZAdd(context, "longurl", &redis.Z{
		Score:  urlCount,
		Member: "https://2a0xaz2.com/randomurl",
	},
	)
	index, _ := getIndexOfLongURL("https://2a0xaz2.com/randomurl")
	if index != int64(urlCount) {
		t.Fatalf(`Error: %v want match for %v`, index, urlCount)
	}
	client.ZPopMax(context, "longurl", 0)
}

func TestGetIndexOfValueFromSortedSetErr(t *testing.T) {
	client.ZAdd(context, "testurlset", &redis.Z{
		Score:  0,
		Member: "firsturl",
	},
		&redis.Z{
			Score:  1,
			Member: "secondurl",
		},
		&redis.Z{
			Score:  2,
			Member: "thirdurl",
		})
	index, _ := getIndexOfValueFromSortedSet("testurlset", "fourthurl")
	if index != -1 {
		t.Fatalf(`Error: "%v" want match for "%v"`, index, -1)
	}
	client.Del(context, "testurlset")
}
