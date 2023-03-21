package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.OpenFile("/tmp/fetchall", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch, f)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, f *os.File) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	f.WriteString(fmt.Sprintf("%.2ff %7d %s\n", secs, nbytes, url))
	ch <- fmt.Sprintf("%.2ff %7d %s", secs, nbytes, url)
}
