// LICENSE sha256: c71d239df91726fc519c6eb72d318ec65820627232b2f796219e87dcf35d0ab4

package main

import "fmt"

func main() {
	fmt.Printf("Hello\n")
	TestSomething(42)
}

// TestSomething is a function.
func TestSomething(i int) {
	fmt.Printf("Hello again\n")
	i++

	m := map[string]bool{}
	// What are k, and v?
	for k, v := range m {
		fmt.Printf("k:%v,v:%v", k, v)
	}
}
