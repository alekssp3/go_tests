package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func ReadDir(path string) []string{
	out := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(path, file.Name())
		if file.IsDir() {
			out = append(out, path +  string(filepath.Separator) + file.Name())
		}
	}
	return out
}

func ReadDirs(path string)  {
	q := []string{path}
	for len(q) > 0 {
		q = append(q, ReadDir(q[0])...)
		q = q[1:]
	}
}

func main() {
	ReadDirs("c:\\Apps\\autohotkey")
}
