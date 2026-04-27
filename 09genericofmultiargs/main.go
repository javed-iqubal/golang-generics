package main

import "fmt"

func printPair[K any, V any](key K, value V) {

	fmt.Println(key, " ", value)
}

func main() {

	printPair(300, "Three hundred")
	printPair("Two thousand", 2000)
}
