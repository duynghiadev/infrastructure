package proxy

import (
	"sync"
	"time"
)

// RealData Real data object (database query)
type RealData struct {
	ID         string
	Content    string
	LastUpdate time.Time
}

func (r *RealData) FetchData() string {
	// Simulate database query delay
	time.Sleep(500 * time.Millisecond)
	return r.Content
}

// CacheProxy Cache proxy
type CacheProxy struct {
	realData *RealData
	cache    map[string]string
	mu       sync.Mutex
}

func NewCacheProxy(id string, content string) *CacheProxy {
	return &CacheProxy{
		realData: &RealData{ID: id, Content: content, LastUpdate: time.Now()},
		cache:    make(map[string]string),
	}
}

func (c *CacheProxy) FetchData() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if data, exists := c.cache[c.realData.ID]; exists {
		return data // Return the cached data directly
	}

	// First access, get the real data and cache it
	data := c.realData.FetchData()
	c.cache[c.realData.ID] = data
	return data
}
