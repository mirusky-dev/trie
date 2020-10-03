package main

import (
	"fmt"
)

func main() {

	t := NewTrie()

	// Insert
	t.Insert("apple")
	t.Insert("app")
	t.Insert("application")
	t.Insert("homework")

	// Search
	fmt.Println(t.Find("app"))   // expected output: true
	fmt.Println(t.Find("apply")) // expected output: false

	// Delete
	t.Delete("app")

	// search after delete
	fmt.Println(">>>Search after delete<<<")
	fmt.Println(t.Find("app"))         // expected output: false
	fmt.Println(t.Find("application")) // expected output: true

}
