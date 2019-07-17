package main

import (
	"fmt"
	"sync"
)

var (
	//创建一个map存爬取的url
	m = make(map[string]int)
	//创建互斥锁
	lock sync.Mutex
	//群组等待，
	wait sync.WaitGroup
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
	// 下面并没有实现上面两种情况：

	defer wait.Done()

	if depth <= 0 {
		return
	}

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		// fmt.Println(err)
		return
	}
	// fmt.Printf("found: %s %q\n", url, body)
	// for _, u := range urls {
	// 	Crawl(u, depth-1, fetcher)
	// }
	// return

	//存入数据的过程是原子操作，中间不可以打断，所以需要加锁
	lock.Lock()
	//如果这个url没有被爬过
	if m[url] == 0 {
		m[url]++
		depth--
		for _, v := range urls {
			wait.Add(1)
			go Crawl(v, depth, fetcher)
		}
	}
	lock.Unlock()
}

func main() {
	wait.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	wait.Wait() //一直等待，直到子进程任务结束

	for i, _ := range m {
		fmt.Println(i) //i就是map中的key，就是url
	}
	fmt.Println("crawl success and done.")
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

/*
Output:

https://golang.org/
https://golang.org/pkg/
https://golang.org/pkg/os/
https://golang.org/pkg/fmt/
crawl success and done.

*/
