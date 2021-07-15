package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"time"
)

func readFile(name string) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("Error!\n>>>\n%v\n<<<\n", err)
	}
	fmt.Println(string(data))
}

func dateToString(date time.Time) string {
	year, month, day := date.Date()
	return fmt.Sprintf("%2v %2v %v", day, int(month), year)
}

func listDir(info fs.FileInfo) string {
	mode := info.Mode()
	date := dateToString(info.ModTime())
	size := info.Size()
	name := info.Name()
	return mode.String() + " " + date + " " + strconv.FormatInt(size, 10) + " " + name
}

func readDir(name string) {
	data, err := ioutil.ReadDir(name)
	if err != nil {
		fmt.Printf("Error!\n>>>\n%v\n<<<\n", err)
	}
	for i := 0; i < len(data); i++ {
		fmt.Printf("%s\n", listDir(data[i]))
	}
}

//func readDirs(dir string) {
//	dirs := append([]string{}, dir)
//	for len(dirs) > 0 {
//		data, _ := ioutil.ReadDir(dirs[0])
//		for i := 0; i < len(data); i++ {
//			fmt.Println(listDir(data[i]))
//			if data[i].IsDir() {
//				dirs = append(dirs, data[i].Name() + string(os.PathSeparator))
//			}
//		}
//		dirs = dirs[1:]
//		fmt.Println(len(dirs), dirs)
//	}
//}

//func readDirs(path string) {
//	files, _ := filepath.Glob(path)
//
//	for _, i := range files {
//		fmt.Println(i)
//	}
//}


func ReadDir(path string) []string {
	tmp := make([]string, 0)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			abs, _ := filepath.Abs(f.Name())
			tmp = append(tmp, abs)
		}
	}
	return tmp
}

func main() {
	r1 := ReadDir("C:\\Users\\a.spiridonychev\\Projects\\go_tests\\read_file\\go1.16.4")
	//r2 := make([]string, 1)
	fmt.Println(r1)
	for _, dir := range r1 {
		fmt.Println(ReadDir(dir))
	}
}
