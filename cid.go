package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

//返回视频文件大小
func getVideoSize(downUrl string) (size int, err error) {
	client := http.Client{}

	videoSizeReq, _ := http.NewRequest("HEAD", downUrl, nil)
	videoSizeReq.Header.Add("User-Agent", User_Agent)
	videoSizeReq.Header.Add("Referer", Referer)

	re, err := client.Do(videoSizeReq)
	if err != nil {
		return -1, err
	}

	size, _ = strconv.Atoi(re.Header["Content-Length"][0])
	//fmt.Println(size)
	return size, err
}

func getDownUrl(cid string) (string, error) {
	url := GenGetAidChildrenParseFun(cid)
	client := http.Client{}
	r, _ := http.NewRequest("GET", url, nil)
	r.Header.Add("User-Agent", User_Agent)

	res, err := client.Do(r)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	//fmt.Println(string(body))
	durl := cidRespBody{}
	err = json.Unmarshal(body, &durl)
	if err != nil {
		return "", err
	}

	downUrl := durl.Durl[0]["url"]
	//将接口转换为字符串
	var downUrlStr string
	switch downUrl.(type) {
	case string:
		downUrlStr = downUrl.(string)
	}
	//fmt.Println(downUrlStr)
	return downUrlStr, nil
}

func download(downUrl, title string) error {
	file, err := os.Create(title + ".mp4")
	if err != nil {
		return err
	}
	defer file.Close()

	client := http.Client{}

	size, err := getVideoSize(downUrl)
	if err != nil {
		return err
	}

	r, _ := http.NewRequest("GET", downUrl, nil)
	r.Header.Add("User-Agent", User_Agent)
	r.Header.Add("Referer", Referer)

	res, err := client.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()


	counter := &writeCounter{
		bar : newBar(size,title),
		videoSize: uint64(size),
		videoName: title,
	}

	//写入文件
	_, err = io.Copy(file, io.TeeReader(res.Body,counter))
	if err != nil {
		return err
	}

	return nil
}
