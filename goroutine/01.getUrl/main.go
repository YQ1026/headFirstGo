package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Page struct {
	Url string
	Size int
	StatusCode int
}

func getWebSize(url string,channel chan Page){
	res,err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

    //Body表示页面的内容
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// 这里是直接获取网页内容 因为直接获取的是byte字节流 这里转换为string类型来可视化。
	// fmt.Println(string(body))

	// 字节切片的大小与页面大小相同,获取响应的长度并打印出来
	// fmt.Printf("url:%v, size:%v\n",url,len(body))
	channel <- Page{Url: url,Size: len(body),StatusCode: res.StatusCode}

}

func main(){

	startTime := time.Now()


	pages := make(chan Page)
	urls :=[]string{
		"https://cn.netxgroup.uk/",
		"https://cms.netxgroup.uk/",
		"https://crm.netxgroup.uk/",
		"https://www.126.com",
	}

	for i:=0;i<1000;i++{
		urls = append(urls,"https://ucloud.cn")
	}

	for _,url :=range urls{
		go getWebSize(url,pages)
	}
	for i:=0;i<(len(urls));i++{
		page := <- pages
		fmt.Printf("url:%v, statusCode:%d, size:%v\n",page.Url,page.StatusCode,page.Size)
	}
	fmt.Println("cost time:",time.Now().Sub(startTime))
}