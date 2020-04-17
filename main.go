package main

import (
	"fmt"
	"os"
)


func main() {

	var bvid string
	fmt.Println("请输入BVID(即b站视频链接BV后的一串代码,例如:1qk4y197bB):")
	_,err := fmt.Scan(&bvid)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取标题,创建文件夹
	allVideos,err = getCids(bvid)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Mkdir(mainTitle,777)
	if err != nil {
		fmt.Println("创建文件夹失败:",err)
		return
	}

	//获取工作路径
	pwd,err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	//更改工作路径到下载目录
	err = os.Chdir(pwd + "/" + mainTitle)
	if err != nil {
		fmt.Println(err)
		return
	}

	wgBar.Add(routines)
	w := &worker{}
	w.work()

	//等待进度条全部到达 100%
	p.Wait()

	//阻塞等待完成
	c := <- result

	fmt.Println(c)
}
