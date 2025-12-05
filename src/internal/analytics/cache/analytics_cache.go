package cache

import (
	"time"

	"github.com/dgraph-io/ristretto/v2"
)

var (
	ttl = 10 * time.Minute
)

type Cache struct {
	versionDownloadCounts *ristretto.Cache[string, int64]
}

func NewCache() *Cache {
	versionDownloadCounts, err := ristretto.NewCache(&ristretto.Config[string, int64]{
		NumCounters: 100 * 10,         // x10 of expected number of elements when full
		MaxCost:     16 * 1024 * 1024, // 16 MB
		BufferItems: 64,               // keep 64
	})
	if err != nil {
		panic(err)
	}

	return &Cache{
		versionDownloadCounts: versionDownloadCounts,
	}
}

func (c *Cache) GetDownloadCountForVersion(spaceID, versionID string) int64 {
	key := spaceID + ":" + versionID

	count, found := c.versionDownloadCounts.Get(key)
	if !found {
		return -1
	}

	return count
}

func (c *Cache) SetDownloadCountForVersion(spaceID, versionID string, count int64) {
	key := spaceID + ":" + versionID

	c.versionDownloadCounts.SetWithTTL(key, count, 4, ttl)
}
