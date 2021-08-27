package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// You should always use Join instead of concatenating /s or \s manually.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p: ", p) // p: dir1/dir2/filename
	
	fmt.Println(filepath.Join("dir1//", "filename")) // dir1/filename
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // dir1/filename

	// Dir and Base can be used to split a path to the directory and the file. Alternatively, Split will return both in the same call.
	fmt.Println("Dir(p):", filepath.Dir(p)) // Dir(p): dir1/dir2
	fmt.Println("Base(p):", filepath.Base(p)) // Base(p): filename

	// check whether a path is absolute
	fmt.Println(filepath.IsAbs("dir/file")) // false
	fmt.Println(filepath.IsAbs("/dir/file")) // true

	filename := "config.json"

	// split the extension of filenames with Ext
	ext := filepath.Ext(filename) 
	fmt.Println(ext) // .json

	// the filename with the extension removed
	fmt.Println(strings.TrimSuffix(filename, ext)) // config

	// Rel finds a relative path between a base and a target. It returns an error if the target cannot be made relative to base
	rel1, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel1) // t/file

	rel2, err := filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel2) // ../c/t/file
}