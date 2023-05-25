package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	info  map[string][]byte
	mutex sync.RWMutex
}

func (c *cache) get(key string) []byte {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.info[key]
}

func (c *cache) add(key string, duration time.Duration, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.info[key] = val
	time.AfterFunc(duration, func() { c.delete(key) })
}

func (c *cache) delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.info, key)
}
