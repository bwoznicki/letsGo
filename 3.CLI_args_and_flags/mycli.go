package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	word        string
	trueorfalse bool
	number      int
)

func main() {

	// set the usage output
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] [resource...]\n\n", os.Args[0])
		fmt.Printf("Example: %s -word someWord -number 10 -trueorfalse \n\n", os.Args[0])

		flag.PrintDefaults()
	}

	flag.StringVar(&word, "word", "default_word", "print some word")
	flag.BoolVar(&trueorfalse, "trueorfalse", false, "bool switch")
	flag.IntVar(&number, "number", 0, "int flag")

	// retrieve the flags
	flag.Parse()

	osargs := os.Args
	args := flag.Args()

	fmt.Printf("os.Args() = %v\n\n", osargs)
	fmt.Printf("flag.Args() = %v\n\n", args)

	fmt.Printf("flag word = %v\n", word)
	fmt.Printf("flag trueorfalse = %v\n", trueorfalse)
	fmt.Printf("flag number = %v\n", number)

}
