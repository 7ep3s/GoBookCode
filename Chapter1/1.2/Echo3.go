package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], "\r\n"))
	nanosecs := time.Since(start).Nanoseconds()
	fmt.Println("%d", nanosecs)
}
