// https://thedevelopercafe.com/articles/singleton-in-golang-839d8610958b

package cache

import "sync"

// Interface our cache needs to fullfill
type MyCache interface {
	Get(key string) string
	Set(key string, value string)
}

// Create a struct that implements the interface
// Struct name starts with lower case letter as we don't want anyone
// to be able to create a new instance of myCache
type myCache struct {
}

// Teh actual implementation of `Get` and `Set` don't matter for this demo
func (c myCache) Get(key string) string {
	return "default value"
}

func (c myCache) Set(key string, value string) {

}

// GetCache function to always return the _same_ instance of myCache
var cache *myCache


func GetCache() MyCache {
	once.Do(func() {
		cache = &myCache{}
	})



