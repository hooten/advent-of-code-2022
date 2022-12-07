package main

import (
	"fmt"
	"github.com/hooten/advent-of-code-2022/pkg/util"
	"github.com/kr/pretty"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func readFile() string {
	bytes, err := os.ReadFile("./internal/dec07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func getLines() []string {
	fileBlob := readFile()
	rawLines := strings.Split(fileBlob, "\n")
	return util.Filter(
		func(s string) bool {
			return s != ""
		},
		rawLines,
	)
}

func getChars() []string {
	fileBlob := readFile()
	rawChars := strings.Split(fileBlob, "")
	return util.Filter(
		func(s string) bool {
			return s != "" && s != "\n"
		},
		rawChars,
	)

}

type Tree struct {
	name   string
	kind   string
	size   int
	files  []*Tree
	parent *Tree
}

func main() {
	lines := getLines()
	filesystem := &Tree{
		name:   "/",
		kind:   "dir",
		size:   0,
		files:  []*Tree{},
		parent: nil,
	}
	fmt.Println(filesystem)
	var position *Tree
	var listing bool
	for _, line := range lines {
		cdCommand := regexp.MustCompile("\\$ cd (.*)")
		cdMatches := cdCommand.FindAllStringSubmatch(line, -1)
		if len(cdMatches) > 0 {
			listing = false
			nextDir := cdMatches[0][1]
			if nextDir == "/" {
				position = filesystem
			}
			if nextDir == ".." {
				position = position.parent
			}
			for _, dir := range position.files {
				if dir.name == nextDir {
					position = dir
					break
				}
			}
			continue
		}
		lsCommand := regexp.MustCompile("\\$ ls")
		lsMatches := lsCommand.FindAllStringSubmatch(line, -1)
		if len(lsMatches) > 0 {
			listing = true
			continue
		}
		if listing {
			dirList := regexp.MustCompile("dir (.*)")
			dirMatches := dirList.FindAllStringSubmatch(line, -1)
			if len(dirMatches) > 0 {
				position.files = append(position.files, &Tree{
					name:   dirMatches[0][1],
					kind:   "dir",
					size:   0,
					files:  []*Tree{},
					parent: position,
				})
			}
			fileList := regexp.MustCompile("(\\d+) (.*)")
			fileMatches := fileList.FindAllStringSubmatch(line, -1)
			if len(fileMatches) > 0 {
				sizeStr := fileMatches[0][1]
				size, errr := strconv.Atoi(sizeStr)
				if errr != nil {
					log.Fatal(errr)
				}
				name := fileMatches[0][2]
				position.files = append(position.files, &Tree{
					name:   name,
					kind:   "file",
					size:   size,
					files:  nil,
					parent: position,
				})
			}
			continue
		}
		fmt.Println("should not get here line", line)
	}
	compute(filesystem)
	free := 70000000 - filesystem.size
	need := 30000000 - free
	pretty.Print(need)
	getDirs(filesystem)
	ints := []int{
		19628923,
		15328762,
		12545514,
	}
	sort.Ints(ints)
	fmt.Println(ints[0])
}

func getDirs(t *Tree) {
	if t.kind == "dir" {
		if t.size >= 10216456 {
			fmt.Println(t.size)
		}
	}
	for _, subtree := range t.files {
		getDirs(subtree)
	}
}

func compute(t *Tree) int {
	if t.files == nil {
		return t.size
	}
	total := 0
	for _, subtree := range t.files {
		total += compute(subtree)
	}
	t.size = total
	return t.size
}