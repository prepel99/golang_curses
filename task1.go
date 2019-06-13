package main

import (
	"flag"
	"fmt"
)

func main() {
	width := flag.Uint64("w", 5, "set width > 0")
	height := flag.Uint64("h", 5, "set height > 0")
	flag.Parse()
  flagSetCount := 0
  flag.Visit(func(f *flag.Flag) {
    flagSetCount++
  })
  if *width >uint64(100) || *height > uint64(100) {
    fmt.Println("There is not correct one of the params")
    panic()
  }
  if flagSetCount == 0 {
    fmt.Println("HELP: there is 2 flags, both must be 0 < dighit<= 100, you can set width with -w, height with -h")
  } else if flagSetCount < 2 {
    fmt.Println("Both of params should be set correct")
  } else {
	 outPut(*width, *height)
  }
}

func outPut(width, height uint64) {
	for i := uint64(0); i < height; i++ {
		if i%2 != 0 {
			fmt.Printf(" ")
		}
		for j := uint64(0); j < width; j++ {
      fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}
