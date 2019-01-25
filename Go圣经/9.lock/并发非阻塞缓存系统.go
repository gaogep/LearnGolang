package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mome"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	urls := []string{"https://www.17173.com/",
		"https://www.qidian.com/",
		"http://nb.zol.com.cn/",
		"https://www.taobao.com/",
		"https://weibo.com/",
		"https://www.icourse163.org/",
		"https://nba.hupu.com/",
		"https://www.bilibili.com/",
		"https://www.17173.com/",
		"https://www.taobao.com/",
		"https://www.qidian.com/"}

	m := memo.New(httpGetBody)
	var n sync.WaitGroup
	t := time.Now()
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s %s %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
	fmt.Println("总时间为:", time.Since(t))
}
