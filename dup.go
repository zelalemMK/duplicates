package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sort"
)

var prereq = map[string][]string{
	"algorithims": {"data structres"},
	"calculus":    {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	filename, _, _ := fetch("https://www.example.com")
	fmt.Println(filename)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.RequestURI)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", 0, err
	}

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err

}

func max(n ...int) int {
	maxVal := 0
	for _, i := range n {
		if i > maxVal {
			maxVal = i
		}
	}
	return maxVal
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
