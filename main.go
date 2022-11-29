package main

import (
	"fmt"
	"lru_cache/LRUCache"
)

func main() {
	c := LRUCache.NewLRUCache(5)

	c.Add("1", "1")
	c.Add("2", "2")
	c.Add("3", "3")
	printC(c)
	c.Get("1")

	printC(c)
	c.Add("4", "4")
	c.Add("5", "5")

	printC(c)
	c.Add("6", "6")
	c.Add("7", "7")
	c.Add("7", "7")

	printC(c)

	c.Remove("ad")
	//c.Remove("6")

	printC(c)
	fmt.Println(c.Get("6"))
	fmt.Println(c.Get("1"))
	printC(c)
}

func printC(c *LRUCache.LRUCache) {
	fmt.Println("______")
	m, q := c.GetAll()
	for k, v := range m {
		fmt.Println("key: ", k, "val:", v)
	}
	cur := q.Head
	fmt.Println(c.Len(), c.Cap())
	for cur != nil {
		fmt.Println("node: ", cur.Value)
		cur = cur.Next
	}
}
