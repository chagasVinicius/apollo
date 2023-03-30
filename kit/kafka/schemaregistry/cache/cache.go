package cache

// Cache represents a key-value storage where to put cached data.
type Cache interface {
	// Get returns the cache value associated with key
	Get(key interface{}) (interface{}, bool)
	// Put puts a value in cache associated with key
	Put(key interface{}, value interface{})
	// Delete deletes the cache entry associated with key
	Delete(key interface{})
	// ToMap returns the current cache entries copied into a map
	ToMap() map[interface{}]interface{}
}
