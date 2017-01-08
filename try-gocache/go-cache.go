package main

import (
	"fmt"
	"time"

	"github.com/k0kubun/pp"
	cache "github.com/patrickmn/go-cache"
)

type MyStruct struct {
	Id   int
	Name string
}

func MyFunction(s string) { pp.Println(s) }

func main() {

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 30 seconds
	c := cache.New(5*time.Minute, 30*time.Second)

	// Set the value of the key "foo" to "bar", with the default expiration time
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set the value of the key "baz" to 42, with no expiration time
	// (the item won't be removed until it is re-set, or removed using
	// c.Delete("baz")
	c.Set("baz", 42, cache.NoExpiration)

	// Get the string associated with the key "foo" from the cache
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	var x interface{}

	// Since Go is statically typed, and cache values can be anything, type
	// assertion is needed when values are being passed to functions that don't
	// take arbitrary types, (i.e. interface{}). The simplest way to do this for
	// values which will only be used once--e.g. for passing to another
	// function--is:

	foo, found = c.Get("foo")
	if found {
		MyFunction(foo.(string))
	}

	// This gets tedious if the value is used several times in the same function.
	// You might do either of the following instead:
	if x, found = c.Get("foo"); found {
		foo = x.(string)
		// ...
	}
	// or
	var fool string
	if x, found = c.Get("foo"); found {
		fool = x.(string)
	}
	pp.Println(fool)
	// ...
	// foo can then be passed around freely as a string

	// Want performance? Store pointers!
	c.Set("foo", &MyStruct{Id: 10, Name: "abc"}, cache.DefaultExpiration)
	if x, found = c.Get("foo"); found {
		foox := x.(*MyStruct)
		// ...
	}

	// If you store a reference type like a pointer, slice, map or channel, you
	// do not need to run Set if you modify the underlying data. The cached
	// reference points to the same memory, so if you modify a struct whose
	// pointer you've stored in the cache, retrieving that pointer with Get will
	// point you to the same data:
	foos := &MyStruct{Id: 1}
	c.Set("foo", foos, cache.DefaultExpiration)
	// ...
	x, _ = c.Get("foo")
	foo2 := x.(*MyStruct)
	fmt.Println(foo2.Id)
	// ...
	foo2.Id++
	// ...
	x, _ = c.Get("foo")
	foo3 := x.(*MyStruct)
	pp.Println(foo3.Id)

	// will print:
	// 1
	// 2

}
