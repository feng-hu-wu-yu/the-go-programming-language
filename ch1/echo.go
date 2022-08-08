package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string = "", ""
	// 1.1
	fmt.Printf("arg0: %s\n", os.Args[0])
	// 1.2
	for idx, arg := range os.Args[0:] {
		fmt.Printf("arg: %d, %s \n\n", idx, arg)
	}
	// 1.3
	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Printf("for args: %s \n", s)
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("for time: %d \n", cost)
	start = time.Now()
	fmt.Printf("join args: %s \n", strings.Join(os.Args[0:], " "))
	end = time.Now()
	cost = end.Sub(start)
	fmt.Printf("join time: %d \n", cost)

}
