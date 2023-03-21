# GoCache

This project provides a simple, memory-efficient in-memory Key-Value cache, which can be easily integrated into your Go projects. The cache is built on top of [BuntDB](https://github.com/tidwall/buntdb), a fast, embeddable, pure-Go Key-Value database.

## Features

* Simple and easy-to-use API
* Efficient memory usage
* Time-to-live (TTL) support for cache entries
* Thread-safe operations

## Installation

To use this cache in your project, simply run:

```sh
go get github.com/chand1012/gocache
```

## Usage

Here's an example of how to use the cache in your code:

```go

package main

import (
	"fmt"
	"time"

	cache "github.com/chand1012/gocache"
)

func main() {
	// Create a new cache instance
	c, err := cache.New()
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Set a value with a 5-second TTL
	err = c.Set("key", "value", 5*time.Second)
	if err != nil {
		panic(err)
	}

	// Retrieve the value from the cache
	value, err := c.Get("key")
	if err != nil {
		panic(err)
	}

	fmt.Println("The value for 'key' is:", value)
}
```

## API

The API for the cache includes the following methods:

* `New() (*Cache, error)`: Creates a new cache instance
* `(c *Cache) Close() error`: Closes the cache instance and releases resources
* `(c *Cache) Set(key, value string, ttl time.Duration) error`: Stores a key-value pair in the cache with the specified TTL
* `(c *Cache) Get(key string) (string, error)`: Retrieves the value associated with the given key from the cache

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.
