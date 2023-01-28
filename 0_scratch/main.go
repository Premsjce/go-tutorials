package main

import (
	"fmt"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Called recursively which crawls upto  maximum of  depth provided
func Crawl(url string, depth int, fetcher Fetcher, res chan string) {
	defer close(res)

	if depth <= 0 {
		return
	}

	body, ursl, err := fetcher.Fetch(url)
	if err != nil {
		res <- err.Error()
		return
	}

	res <- fmt.Sprintf("Found : %v, %q\n", url, body)

	result := make([]chan string, len(ursl))
	for i, urlItem := range ursl {
		result[i] = make(chan string)
		go Crawl(urlItem, depth-1, fetcher, result[i])
	}

	for i := range result {
		for s := range result[i] {
			res <- s
		}
	}

	return
}

func (f fakeFetcher) Fetch(url string) (body string, urls []string, err error) {

	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}

	return "", nil, fmt.Errorf("URL is not found, %s\n", url)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

var fetcherData = fakeFetcher{
	"https://golang.org/": &fakeResult{
		body: "The Go Programming Language",
		urls: []string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		body: "Packages",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		body: "Package fmt",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		body: "Package os",
		urls: []string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func main() {
	result := make(chan string)
	go Crawl("https://golang.org/", 4, fetcherData, result)

	for s := range result {
		fmt.Println(s)
	}
}
