// goCat.go a simple Go version of cat program

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// check arguments passed to the program
	if len(os.Args) < 2 {
		fmt.Println("Usage: gocat <file1> <file2> ...")
	}

	// in the loop, for each file passed
	// read the file content and print to std.out
	files := os.Args[1:]
	for _, filename := range files {

		//fmt.Println(filename)

		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("os.Open error:\n%s\n", err)
			os.Exit(1)
		}
		defer file.Close()

		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("file.Read error:\n%s\n", err)
			os.Exit(1)
		}

		fmt.Println(string(content))
	}
}
