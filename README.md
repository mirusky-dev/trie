# Trie

Simple trie Insert/Find/Delete implementation using Go

## Installation

This project uses Go Modules, so you could install running the following:

```bash
go get github.com/mirusky-dev/trie
```

## Usage

```go
package main

import (
	"fmt"
)

func main() {

	// Create a new Trie
	t := NewTrie() 

	// Insert items
	t.Insert("apple")
	t.Insert("app")
	t.Insert("application")
	t.Insert("homework")

	// Search for items
	fmt.Println(t.Find("app"))   // expected output: true
	fmt.Println(t.Find("apply")) // expected output: false

	// Delete item
	t.Delete("app")

	// Searching after delete
	fmt.Println(t.Find("app"))         // expected output: false
	fmt.Println(t.Find("application")) // expected output: true

}

```

## License
[MIT](https://choosealicense.com/licenses/mit/)
