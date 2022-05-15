package main

import (
	"fmt"
	"<module_path/cache"
)

func main() {
	c := cache.GetCache()

	fmt.Println(c.Get("key"))
}
