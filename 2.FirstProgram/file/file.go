package file

import (
	"io/ioutil"
	"os"
)

func Read(filename string) (string, error) {

	var content string

	f, err := os.Open(filename)
	if err != nil {
		return content, err
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return content, err
	}

	content = string(c)

	return content, nil
}
