// goCat.go a simple Go version of cat program

package main

import (
	"fmt"
	"os"

	"./file"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: gocat <file1> <file2> ...")
	}

	files := os.Args[1:]
	for _, filename := range files {

		content, err := file.Read(filename)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(content)

	}

}
