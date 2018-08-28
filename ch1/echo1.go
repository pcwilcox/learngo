// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	ch := make(chan string)
	go primitive(ch)
	go advanced(ch)
	for i := 0; i < 2; i++ {
		fmt.Println(<-ch)
	}
}

func primitive(ch chan<- string) {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	stop := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%s | Primitive done in %3f seconds", s, stop)
}

func advanced(ch chan<- string) {
	start := time.Now()
	s := strings.Join(os.Args, " ")
	stop := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%s | Advanced done in %3f seconds", s, stop)
}
