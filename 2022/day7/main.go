package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
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

func NewFile(name string, size int) *File {
	return &File{
		Name: name,
		Size: size,
	}
}

func (d *Dir) AddDir(name string) {
	d.Dirs = append(d.Dirs, NewDir(name, d))
}

func (d *Dir) AddFile(name string, size int) {
	d.Files = append(d.Files, NewFile(name, size))
}

func (d *Dir) Size() int {
	size := 0
	for _, file := range d.Files {
		size += file.Size
	}

	for _, dir := range d.Dirs {
		size += dir.Size()
	}

	return size
}

func SprintLine(line string, tabs int) string {
	var sb strings.Builder

	// print 2 spaces per tab
	for i := 0; i < tabs; i++ {
		sb.WriteString("  ")
	}

	sb.WriteString(fmt.Sprintf("- %s\n", line))
	return sb.String()
}

func (d *Dir) PrintTree(tabs int) string {
	var sb strings.Builder

	line := fmt.Sprintf("%s (dir)", d.Name)
	sb.WriteString(SprintLine(line, tabs))

	for _, dir := range d.Dirs {
		sb.WriteString(dir.PrintTree(tabs + 1))
	}

	for _, file := range d.Files {
		line := fmt.Sprintf("%s (file, size=%v)", file.Name, file.Size)
		sb.WriteString(SprintLine(line, tabs+1))
	}

	return sb.String()
}

type Accumulator func(Dir, int) int

func (d *Dir) Accumulate(f Accumulator, acc int) int {
	acc = f(*d, acc)

	if len(d.Dirs) == 0 {
		// leaf
		return acc
	}

	for _, dir := range d.Dirs {
		acc = dir.Accumulate(f, acc)
	}
	return acc
}

func (d *Dir) Traverse(f func(Dir) bool) {
	shouldContinue := f(*d)

	if !shouldContinue {
		return
	}

	for _, dir := range d.Dirs {
		dir.Traverse(f)
	}
}

type Terminal struct {
	rootDir    *Dir
	currentDir *Dir
}

func NewTerminal() Terminal {
	root := NewDir("/", nil)
	return Terminal{
		rootDir:    root,
		currentDir: root,
	}
}

func (t *Terminal) Cd(dirName string) error {
	log.Println("cd", dirName)

	if dirName == "/" {
		t.currentDir = t.rootDir
		return nil
	}

	if dirName == ".." {
		if t.currentDir == nil {
			return errors.New("no current directory")
		}
		if t.currentDir.Parent == nil {
			return errors.New("can not cd .. from root")
		}
		t.currentDir = t.currentDir.Parent
		return nil
	}

	// descendent dir
	for _, childDir := range t.currentDir.Dirs {
		if childDir.Name == dirName {
			t.currentDir = childDir
			return nil
		}
	}

	return fmt.Errorf("could not find dir %s", dirName)
}

func (t *Terminal) Ls() {
	fmt.Print(t.rootDir.PrintTree(0))
}

func (t *Terminal) Mkdir(dirName string) error {
	log.Println("mkdir", dirName)
	if t.currentDir == nil {
		return errors.New("no current directory")
	}
	t.currentDir.AddDir(dirName)
	return nil
}

func (t *Terminal) Touch(name string, size int) error {
	log.Println("touch", name, size)
	if t.currentDir == nil {
		return errors.New("no current directory")
	}

	t.currentDir.AddFile(name, size)
	return nil
}

func (t *Terminal) Root() *Dir {
	return t.rootDir
}

func main() {
	// read input
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	t := NewTerminal()

	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(line, "$ cd") {
			dirName := line[5:]
			if err := t.Cd(dirName); err != nil {
				log.Fatal(err)
			}
		} else if strings.HasPrefix(line, "dir") {
			dirName := line[4:]
			t.Mkdir(dirName)
		} else if line == "$ ls" {
			// do nothing
		} else {
			split := strings.Split(line, " ")
			if len(split) < 2 {
				log.Fatalf("bad line %s", line)
			}
			name := split[1]
			size, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}
			t.Touch(name, size)
		}
	}

	t.Ls()

	part1 := t.Root().Accumulate(func(d Dir, acc int) int {
		size := d.Size()
		if size < 100000 {
			return acc + size
		}
		return acc
	}, 0)
	log.Println("part1", part1)

	part2Candidates := []Dir{}

	updateSize := 30000000
	freeSpace := 70000000 - t.rootDir.Size()
	targetSize := updateSize - freeSpace

	t.Root().Traverse(func(d Dir) bool {
		if d.Size() >= targetSize {
			part2Candidates = append(part2Candidates, d)
		}
		return true
	})

	sort.Slice(part2Candidates, func(i, j int) bool {
		return part2Candidates[i].Size() < part2Candidates[j].Size()
	})

	log.Println("part2", part2Candidates[0].Name, part2Candidates[0].Size())
}
