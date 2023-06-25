package localcache

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/patrickmn/go-cache"
)

var local ILocalCache

type ILocalCache interface {
	Save(key string, value interface{})
	SaveWithExpiration(key string, value interface{}, t time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
	Increment(key string, n int)
	DeleteSet(key string)
}

type localCache struct {
	c *cache.Cache
}

func NewCache() {
	defaultSecStr := os.Getenv("CACHE_DEFAULT_EXPIRATION_SEC")
	sec, err := strconv.Atoi(defaultSecStr)
	if err != nil {
		panic(err)
	}
	local = &localCache{c: cache.New(time.Duration(sec)*time.Second, 10*time.Second)}
}

func Cache() ILocalCache {
	return local
}

func (lc *localCache) SaveWithExpiration(key string, value interface{}, t time.Duration) {
	lc.c.Set(key, value, t)
}

func (lc *localCache) Save(key string, value interface{}) {
	lc.c.SetDefault(key, value)
}

func (lc *localCache) Get(key string) (interface{}, bool) {
	value, existed := lc.c.Get(key)
	return value, existed
}

func (lc *localCache) Delete(key string) {

	lc.c.Delete(key)
}

func (lc *localCache) DeleteSet(key string) {
	cacheSet := lc.c.Items()

	for k := range cacheSet {
		if strings.Contains(k, key) {
			lc.Delete(k)
		}
	}
}

func (lc *localCache) Increment(key string, n int) {
	lc.c.IncrementInt(key, n)
}
