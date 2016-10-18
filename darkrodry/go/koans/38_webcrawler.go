package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type UrlCache struct {
	cache map[string]bool
	mux   sync.Mutex
}

var urlCache UrlCache

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, finish chan bool) {
	urlCache.mux.Lock()
	urlCache.cache[url] = true
	urlCache.mux.Unlock()
	
	if depth <= 0 {
		finish <- true
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		finish <- false
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	done := make(chan bool)
	toCrawl := 0
	for _, u := range urls {
		urlCache.mux.Lock()
		if _, ok := urlCache.cache[u]; !ok {
			toCrawl = toCrawl + 1
			go Crawl(u, depth-1, fetcher, done)
		}
		urlCache.mux.Unlock()
	}

	for i := 0; i < toCrawl; i++ {
		<-done
	}

	finish <- true
	return
}

func main() {
	urlCache = UrlCache{cache: make(map[string]bool)}
	done := make(chan bool)
	go Crawl("http://golang.org/", 4, fetcher, done)
	<-done
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
