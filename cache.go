package gocache

import (
	"time"

	"github.com/tidwall/buntdb"
)

// Cache fully in memory cache
type Cache struct {
	db *buntdb.DB
}

// New creates a new cache
func New() (*Cache, error) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		return nil, err
	}

	return &Cache{db: db}, nil
}

// Close closes the cache
func (c *Cache) Close() error {
	return c.db.Close()
}

// Set sets a key/value pair in the cache
func (c *Cache) Set(key, value string, ttl time.Duration) error {
	return c.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, &buntdb.SetOptions{Expires: true, TTL: ttl})
		return err
	})
}

// Get gets a value from the cache
func (c *Cache) Get(key string) (string, error) {
	var value string
	err := c.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		value = val
		return nil
	})
	return value, err
}
