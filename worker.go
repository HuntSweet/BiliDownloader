package main

import (
	"fmt"
	"strconv"
	"sync"
)

type worker struct{

}

func (w *worker) work() {
	downQueue := make(chan video,len(allVideos))
	for _,v := range allVideos{
		downQueue <- v
	}

	wg := &sync.WaitGroup{}

	wg.Add(routines)

	for i:=0;i<routines;i++{
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			defer wgBar.Done()
			for{
				m,ok := <- downQueue
				if !ok{
					return
				}
				downUrl,err := getDownUrl(strconv.Itoa(m.cid))
				if err != nil{
					fmt.Println("getDownUrl:",err)
					return
				}
				//fmt.Println(m.title)
				err = download(downUrl,m.title)
				if err != nil{
					fmt.Println(m.title + "下载失败")
				}
			}

		}(wg)
	}

	close(downQueue)
	wg.Wait()

	//如果只声明，那么就它就是一个nil通道，无法往里面写值
	result <- "All done"

}