package main

import (
	"flag"
	"fmt"
)

func main() {
	width := flag.Uint64("w", 0, "set width > 0")
	height := flag.Uint64("h", 0, "set height > 0")
	flag.Parse()
  flagSetCount := 0
  flag.Visit(func(f *flag.Flag) {
    flagSetCount++
  })
  if flagSetCount < 2 {
    fmt.Println("Params should be set")
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
