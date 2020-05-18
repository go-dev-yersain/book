package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

// Lookup returns cache's value by provided key
func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

// Rocket is the prototype of the best rocket in the world
type Rocket struct{}

// Launch is the best start plant for the best rocket
func (r *Rocket) Launch() {
	fmt.Println("The rocket has been launched")
}

func main() {
	// p, q := geometry.Point{1, 2}, geometry.Point{4, 6}
	// distance := geometry.Point.Distance
	// fmt.Println(distance(p, q))
	// fmt.Printf("%T\n", distance)

	// scale := (*geometry.Point).ScaleBy
	// scale(&p, 2)
	// scale(&p, 2)
	// fmt.Println(p)
	// fmt.Printf("%T\n", scale)

	// paths := geometry.Path{{0, 0}, {1, 2}, {3, 4}, {4, 5}}
	// fmt.Println(paths)
	// translateBy := (geometry.Path).TranslateBy
	// translateBy(paths, geometry.Point{4, 4}, true)
	// fmt.Println(paths)

	// var x, y intset.IntSet
	// x.Add(1)
	// x.Add(144)
	// x.Add(9)
	// fmt.Println(x.String())

	// y.Add(9)
	// y.Add(42)
	// fmt.Println(y.String())

	// x.UnionWith(&y)
	// fmt.Println(x.String())

	// fmt.Println(&x)
	// fmt.Println(x.String())
	// fmt.Println(x)

}
