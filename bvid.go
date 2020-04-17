package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type PageList struct {
	Data pageListData `json:"data"`
}

type pageListData struct {
	Title string `json:"title"`
	Pages []map[string]interface{} `json:"pages"`
}

func getCids(bvid string) (videos []video,err error) {
	cidsUrl := fmt.Sprintf("%sbvid=%s",videoInfoUrl,bvid)

	client := http.Client{}
	r,err := http.NewRequest("GET",cidsUrl,nil)
	if err != nil{
		return videos,err
	}

	r.Header.Add("User-Agent", User_Agent)
	resp ,err := client.Do(r)
	if err != nil{
		return videos,err
	}
	defer resp.Body.Close()

	re,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return videos,err
	}
	//fmt.Println(string(re))

	pages := PageList{}
	err = json.Unmarshal(re,&pages)
	//fmt.Println(pages.Data.Pages)
	mainTitle = pages.Data.Title
	//fmt.Println(mainTitle)
	for _,v := range pages.Data.Pages{
		temp := video{}
		//默认是float64类型，转换为int之后在转换为字符串
		page := strconv.Itoa(int(v["page"].(float64)))
		//视频分P
		temp.title = "P" + page + " " + v["part"].(string)
		temp.cid = int(v["cid"].(float64))
		videos = append(videos,temp)
	}
	//fmt.Println(videos)
	return videos,nil
}
