// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("Index: %d Arg: %s\n", i, os.Args[i])
	}
}
