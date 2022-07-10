package main

import (
	"bufio"
	"fmt"
	"github.com/iralance/go-clawler/config"
	"github.com/iralance/go-clawler/engine"
	"github.com/iralance/go-clawler/persist"
	"github.com/iralance/go-clawler/scheduler"
	"github.com/iralance/go-clawler/xcar/parser"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 2,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:    "https://newcar.xcar.com.cn/",
		Parser: engine.NewFuncParser(parser.ParseCarList, config.ParseCarList),
	})
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:    "https://newcar.xcar.com.cn/",
	//	Parser: engine.NewFuncParser(parser.ParseCarList, config.ParseCarList),
	//})
}

func demo() {
	resp, err := http.Get("https://newcar.xcar.com.cn/")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	printBrandList(all)
}

func printBrandList(content []byte) {
	re := regexp.MustCompile(`<a href="(/car/[\d+-]+/)" title="([^"]+)" data-id="(\d+)">`)
	matches := re.FindAllSubmatch(content, -1)
	fmt.Println(matches)
	n := 0
	for _, match := range matches {
		n++
		fmt.Printf("href: %s, title: %s, id: %s", match[1], match[2], match[3])
		fmt.Println()
	}

	fmt.Println("all count ", n)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
