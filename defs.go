package main

import (
	mpb "github.com/vbauerster/mpb/v5"
	"sync"
)

var (
	//获取b站视频cid列表
	pageListUrl = "https://api.bilibili.com/x/player/pagelist?"
	//视频信息
	videoInfoUrl = "https://api.bilibili.com/x/web-interface/view?"
	cidApiUrl    = "https://api.bilibili.com/x/player/playurl?"
	User_Agent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.13; rv:56.0) Gecko/20100101 Firefox/56.0"
	Referer      = "https://www.bilibili.com"
	//主标题
	mainTitle string
	//所有的视频
	allVideos []video
	//协程数目
	routines = 10
	//是否已经结束,必须是有缓存的通道
	result = make(chan string,1)

	wgBar = &sync.WaitGroup{}
	p = mpb.New(mpb.WithWaitGroup(wgBar))

	mu sync.Mutex
)

//type muteP struct {
//	mu sync.RWMutex
//	p *mpb.Progress
//}

type video struct {
	title string
	cid int
}

type cidRespBody struct {
	Durl []map[string]interface{} `json:"durl"`
}


type writeCounter struct {
	Total uint64
	videoSize uint64
	videoName string
	bar *mpb.Bar
}

//必须使用指针，否则只会返回每次下载字节数，而不是总字节数
func (w *writeCounter) Write(p []byte) (int,error) {
	n := len(p)
	w.Total += uint64(n)

	w.printProgress(n)
	return n,nil

}

func (w *writeCounter) printProgress(n int)  {
	w.bar.IncrBy(n)
	//percent := float64(w.Total) / float64(w.videoSize) * 100
	//fmt.Printf("\r %s 已下载 %.2f %s \n",w.videoName,percent,"%")
}
