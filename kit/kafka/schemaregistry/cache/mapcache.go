package cache

// MapCache is a cache backed by a map.
type MapCache struct {
	entries map[interface{}]interface{}
}

// NewMapCache creates a new cache backed by a map.
func NewMapCache() *MapCache {
	c := new(MapCache)
	c.entries = make(map[interface{}]interface{})
	return c
}

// Get returns the cache value associated with key.
func (c *MapCache) Get(key interface{}) (value interface{}, ok bool) {
	value, ok = c.entries[key]
	return
}

// Put puts a value in cache associated with key.
func (c *MapCache) Put(key interface{}, value interface{}) {
	c.entries[key] = value
}

// Delete deletes the cache entry associated with key.
func (c *MapCache) Delete(key interface{}) {
	delete(c.entries, key)
}

// ToMap returns the current cache entries copied into a map.
func (c *MapCache) ToMap() map[interface{}]interface{} {
	ret := make(map[interface{}]interface{})
	for k, v := range c.entries {
		ret[k] = v
	}
	return ret
}
