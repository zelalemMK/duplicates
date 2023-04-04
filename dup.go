package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileWithDup struct {
	FileName string
	Counts   map[string]int
}

func main() {
	var counts []FileWithDup
	files := os.Args[1:]

	if len(files) == 0 {
		// noFileName := FileWithDup{FileName: ""}
		// countLinesStdin(os.Stdin, noFileName.Counts)
		fmt.Println("No file input")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dub2: %v\n", err)
				continue
			}
			countLines(files, &counts)
			f.Close()
		}
	}

	for _, count := range counts {
		showFileName := false
		for line, n := range count.Counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
				showFileName = true
			}
		}
		if showFileName {
			fmt.Println(count.FileName)
		}

	}
}

func countLinesStdin(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func countLines(filenames []string, c *[]FileWithDup) {
	for _, arg := range filenames {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dub2: %v\n", err)
			continue
		}

		d := FileWithDup{
			FileName: arg,
			Counts:   make(map[string]int),
		}

		input := bufio.NewScanner(f)
		for input.Scan() {
			d.Counts[input.Text()]++
		}
		*c = append(*c, d)
	}

}
