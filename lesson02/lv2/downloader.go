package main

import (
	"fmt"
	"sync"
	"time"
)

type Downloader struct {
	FileName string
	wg       *sync.WaitGroup
	result   chan string
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
	for _, fileMame := range fileList {
		d.wg.Add(1)
		go d.download(fileMame)
	}

	go func() {
		d.wg.Wait()
		close(d.result)
	}()
}
