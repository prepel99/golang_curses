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
	if flagSetCount == 0 {
		fmt.Println("HELP: you should set 2 params: -w (width) and -h (height), it must be 0 < dighit <=100 ")
	} else if flagSetCount != 2 {
		fmt.Println("2 params should be set")
	} else if *width <= 100 && *height <= 100 {
		outPut(*width, *height)
	} else {
		fmt.Println("Params are not correct!")
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
