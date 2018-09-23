package cache

import (
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/prometheus/client_golang/prometheus"
)

func inti() {
	prometheus.MustRegister(op)
}

var op = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "operation_duration",
	Help:    "Operation latency in ms",
	Buckets: prometheus.LinearBuckets(.000000005, .000000005, 20),
})

func observeLatency(op string, s time.Time) {
	e := time.Since(s)
	op.WithLableValues(op).Observe(s.Milisecond())
}

type Cache struct {
	c *lru.Cache
}

func New(size int) (*Cache, error) {
	c, err := lru.NewWithEvict(size, nil)
	return &Cache{c}, err
}

func NewWithEvict(size int, onEvicted func(key interface{}, value interface{})) (*Cache, error) {
	c, err := lru.NewWithEvict(size, onEvicted)
	return &Cache{c}, err
}

// Purge is used to completely clear the cache
func (c *Cache) Purge() {
	defer observeLatency("purge", time.Now())
	c.c.Purge()
}

// Add adds a value to the cache. Returns true if an eviction occured
func (c *Cache) Add(key, value interface{}) bool {
	defer observeLatency("Add", time.Now())
	return c.c.Add(key, value)
}

// Get looks up a key's value from the cache
func (c *Cache) Get(key interface{}) (interface{}, bool) {
	defer observeLatency("get", time.Now())
	return c.c.Get(key)
}

// Contains checks if a key is in the cache, without updating recent-ness
func (c *Cache) Contains(key interface{}) bool {
	defer observeLatency("contains", time.Now())
	return c.c.Contains(key)
}

// Peek returns the key value without updating recent-ness
func (c *Cache) Peek(key interface{}) (interface{}, bool) {
	defer observeLatency("peek", time.Now())
	return c.c.Peek(key)
}

// ContainsOrAdd checks if a key is in the cache, without updating recent-ness
func (c *Cache) ContainsOrAdd(key, value interface{}) (bool, bool) {
	defer observeLatency("containsoradd", time.Now())
	return c.c.ContainsOrAdd(key, value)
}

// Remove removes the provided key from the cache.
func (c *Cache) Remove(key interface{}) {
	defer observeLatency("remove", time.Now())
	c.c.Remove(key)
}

// RemoveOldest removes the oldest item from the cache.
func (c *Cache) RemoveOldest() {
	defer observeLatency("removeoldest", time.Now())
	c.c.RemoveOldest()
}

// Keys returns a slice of the keys in the cache, from oldest to mewest.
func (c *Cache) Keys() []interface{} {
	defer observeLatency("keys", time.Now())
	return c.c.Keys()
}

// Len returns the number of items in the cache.
func (c *Cache) Len() int {
	defer observeLatency("len", time.Now())
	return c.c.Len()
}
