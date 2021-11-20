package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

// By is the type of "less" function that defines the ordering of its entry arguments.
type By func(d1, d2 *os.FileInfo) bool

// Sort is a method on the function type, By, that sorts the argument slice according to the function.
func (by By) Sort(dirs []os.FileInfo) {
	ds := &dirSorter{
		entries: dirs,
		by:      by,
	}
	sort.Sort(ds)
}

// dirSorter struct joins our entries and function by
type dirSorter struct {
	entries []os.FileInfo
	by      func(d1, d2 *os.FileInfo) bool
}

// Len is part of sort.Interface.
func (f *dirSorter) Len() int {
	return len(f.entries)
}

// Swap is part of sort.Interface.
func (f *dirSorter) Swap(i, j int) {
	f.entries[i], f.entries[j] = f.entries[j], f.entries[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (f *dirSorter) Less(i, j int) bool {
	return f.by(&f.entries[i], &f.entries[j])
}

func makeDirEntries(entries []os.FileInfo) []os.FileInfo {
	returnable := make([]os.FileInfo, 0)
	for _, file := range entries {
		if file.IsDir() == true {
			returnable = append(returnable, file)
		}
	}
	return returnable
}

func recursion(out io.Writer, entries []os.FileInfo, pF bool, tabs int, path string, isLast []bool) error {
	if pF == false {
		entries = makeDirEntries(entries)
	}
	name := func(d1, d2 *os.FileInfo) bool {
		return (*d1).Name() < (*d2).Name()
	}
	By(name).Sort(entries)
	for index, file := range entries {
		if file.IsDir() == true {
			newEntries, err := ioutil.ReadDir(path + "/" + file.Name())
			if err != nil {
				return fmt.Errorf("Couldnt open file!\n")
			}
			for i := 0; i < tabs; i++ {
				if isLast[i] == true {
					fmt.Fprintf(out, "\t")
				} else {
					fmt.Fprintf(out, "│\t")
				}
			}
			if len(entries)-1 == index {
				fmt.Fprintf(out, "└───")
			} else {
				fmt.Fprintf(out, "├───")
			}
			fmt.Fprintf(out, file.Name()+"\n")
			if len(entries)-1 == index {
				isLast[len(isLast)-1] = true
				err = recursion(out, newEntries, pF, tabs+1, path+"/"+file.Name(), append(isLast, true))
			} else {
				err = recursion(out, newEntries, pF, tabs+1, path+"/"+file.Name(), append(isLast, false))
			}
			if err != nil {
				return fmt.Errorf("Error in recursion!\n")
			}
		} else if file.IsDir() == false && pF == true {
			for i := 0; i < tabs; i++ {
				if isLast[i] == true {
					fmt.Fprintf(out, "\t")
				} else {
					fmt.Fprintf(out, "│\t")
				}
			}
			if len(entries)-1 == index {
				fmt.Fprintf(out, "└───")
			} else {
				fmt.Fprintf(out, "├───")
			}
			if file.Size() == 0 {
				fmt.Fprintf(out, file.Name()+" (empty)"+"\n")
			} else {
				size := strconv.Itoa(int(file.Size()))
				fmt.Fprintf(out, file.Name()+" ("+size+"b)\n")
			}
		}
	}
	return nil
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	dirnames, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("Couldnt open file!\n")
	}
	isLast := make([]bool, 0)
	err = recursion(out, dirnames, printFiles, 0, path, append(isLast, false))
	if err != nil {
		return fmt.Errorf("Error in recursion!\n")
	}
	return nil
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
