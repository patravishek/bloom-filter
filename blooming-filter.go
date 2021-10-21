package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
)

func main() {
	fmt.Println("Enter the string you want to hash")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		line := scanner.Text()
		fmt.Println(hash(line))
	}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
