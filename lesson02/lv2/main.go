package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	res := make(chan string)

	fileList := []string{"file1.zip", "file2.pdf", "file3.mp4"}
	d := initDownloader(wg, res)

	d.startDownloaders(fileList)

	fmt.Printf("开始下载三个文件...\n")

	for msg := range res {
		fmt.Printf("%v \n", msg)
	}

	fmt.Printf("所有文件下载完成!\n")
}
