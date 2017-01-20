package lesson3

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func FileReader() {
	// openn file
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	// close file
	defer file.Close()
	//file ststistic
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func ReadFileWay2() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

func WriteFile() {
	file, err := os.Create("test2.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("Test text for write \n second row")
}

func ReadCatalogue() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	fileInfo, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fl := range fileInfo {
		fmt.Println("->", fl.Name())
	}
}

func WalkInCatalog() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println("#", path)
		return nil
	})
}
