package main

import (
	"fmt"
	"sync"
	"time"
)

type Downloader struct {
	wg     *sync.WaitGroup
	result chan string
}

func initDownloader(wg *sync.WaitGroup, res chan string) *Downloader {
	return &Downloader{
		wg:     wg,
		result: res,
	}
}

func (d *Downloader) download(fileName string) {
	time.Sleep(time.Second * 5)
	d.result <- fmt.Sprintf("%v 下载完成！", fileName)
	d.wg.Done()
}

func (d *Downloader) startDownloaders(fileList []string) {
	for _, fileName := range fileList {
		d.wg.Add(1)
		go d.download(fileName)
	}

	go func() {
		d.wg.Wait()
		close(d.result)
	}()
}
