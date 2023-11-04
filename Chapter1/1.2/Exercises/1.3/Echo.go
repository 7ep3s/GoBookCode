package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	stringsJoinEcho()
	loopEcho()
}

func stringsJoinEcho() {
	start := time.Now()
	asd := strings.Join(os.Args[1:], "\r\n")
	asd = ""
	fmt.Println(asd)
	nanosecs := time.Since(start).Nanoseconds()
	fmt.Println("join speed = %d", nanosecs)
}

func loopEcho() {
	start2 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "\r\n"
	}
	nanosecs2 := time.Since(start2).Nanoseconds()
	s = ""
	fmt.Println(s)
	fmt.Println("loop speed = %d", nanosecs2)
}
