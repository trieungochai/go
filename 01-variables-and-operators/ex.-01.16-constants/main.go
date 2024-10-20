// In this exercise, we have a performance problem: our database server is too slow.
// We are going to create a custom memory cache.
// We’ll use Go’s map collection type, which will act as the cache.
// There is a global limit on the number of items that can be in the cache.
// We’ll use one map to help keep track of the number of items in the cache.
// We have 2 types of data we need to cache: books and CDs.
// Both use the ID, so we need a way to separate the two types of items in the shared cache.
// We need a way to set and get items from the cache.
// We’re going to set the maximum number of items in the cache.
// We’ll also use constants to add a prefix to differentiate between books and CDs.

package main

import "fmt"

const GlobalLimit = 100
const MaxCacheSize int = 10 * GlobalLimit

const (
	CacheKeyBook = "book_"
	CacheKeyCD   = "cd_"
)

var cache map[string]string

func cacheSet(key string, val string) {
	if len(cache)+1 >= MaxCacheSize {
		return
	}

	cache[key] = val
}

func SetBook(isbn string, title string) {
	cacheSet(CacheKeyBook+isbn, title)
}

func SetCD(sku string, title string) {
	cacheSet(CacheKeyCD+sku, title)
}

func cacheGet(key string) string {
	return cache[key]
}

func GetBook(isbn string) string {
	return cacheGet(CacheKeyBook + isbn)
}

func GetCD(sku string) string {
	return cacheGet(CacheKeyCD + sku)
}

func main() {
	cache = make(map[string]string)

	SetBook("1234-5678", "Brave New World")
	SetCD("1234-5678", "Brave New World (Audiobook)")

	fmt.Println("Book :", GetBook("1234-5678"))
	fmt.Println("CD   :", GetCD("1234-5678"))
}
