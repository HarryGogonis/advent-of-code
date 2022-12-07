package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type File struct {
	Name string
	Size int64
}

type Dir struct {
	Name   string
	Parent *Dir
	Files  []*File
	Dirs   []*Dir
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{
		Name:   name,
		Parent: parent,
		Files:  []*File{},
		Dirs:   []*Dir{},
	}
}

func NewFile(name string, size int64) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

func (d *Dir) AddDir(name string) {
	d.Dirs = append(d.Dirs, NewDir(name, d))
}

func (d *Dir) AddFile(name string, size int64) {
	d.Files = append(d.Files, NewFile(name, size))
}

func main() {
	// read input
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	root := NewDir("/", nil)
	var currentDir *Dir

	for i, line := range strings.Split(string(content), "\n") {
		fmt.Printf("%v\t%v\n", i, line)
		if line == "$ cd /" {
			currentDir = root
		} else if line == "$ cd .." {
			if currentDir.Parent == nil {
				panic("cannot cd .. from root")
			}
			currentDir = currentDir.Parent
		} else if strings.HasPrefix(line, "dir") {
			dirName := line[4:]
			fmt.Println("create dir", dirName)
			currentDir.AddDir(dirName)
		} else {
			fmt.Println("not implemented", line)
		}
	}

	fmt.Printf("%+v", root)
}
