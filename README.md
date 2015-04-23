go-cache
========
[![GoDoc](https://godoc.org/github.com/maemual/go-cache?status.svg)](https://godoc.org/github.com/maemual/go-cache)

An in-memory K/V cache and LRU cache library. Just for fun!

This package provide a simple memory key-value cache and LRU cache.
It based on the k/v cache implementation in [go-cache](https://github.com/pmylund/go-cache)
and LRU cache implementation in [groupcache](https://github.com/golang/groupcache/tree/master/lru).

## Documentation

[API Reference](http://godoc.org/github.com/maemual/go-cache)

## Installation

Install go-cache using the "go get" command:

> go get github.com/maemual/go-cache

## Example

Key-value cache:

```
package main

import (
    "fmt"

    "github.com/maemual/go-cache"
)

func main() {
    c := cache.New(0, 0)
    c.Set("1", 1111, 0)
    val, found := c.Get("1")
    if found {
        fmt.Println(val)
    }
    c.Increment("1", 1)
    val, found = c.Get("1")
    if found {
        fmt.Println(val)
    }
}
```

LRU Cache:

```
package main

import (
    "fmt"

    "github.com/maemual/go-cache"
)

func main() {
    lru, err := cache.NewLRU(3)
    if err != nil {
        fmt.Println(err)
    }
    lru.Add("1", 1111)
    lru.Add("2", 2222)
    lru.Add("3", 3333)
    lru.Add("4", 4444)
    _, hit := lru.Get("1")
    if hit {
        fmt.Println("Hit key 1")
    } else {
        fmt.Println("Not hit key 1")
    }

}
```
