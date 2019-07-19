package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

func main() {

	// bookMerge, _ := ioutil.ReadFile("./src/hm.json") 编译前用
	bookMerge, _ := ioutil.ReadFile("hm.json")
	lists := gjson.Get(string(bookMerge), "list")
	time := time.Now().Format("2006-01-02 15:04:05")
	nums := 0
	file := "|书源|API|类型|排行榜|账号|版本|作者|暗码|\n| :----------: | ----------------------------------------- | :---: | :----: | :---: | :---: | :---: |:---: |"
	println(file)
	for _, lis := range lists.Array() {
		author := gjson.Get(lis.String(), "author").String()
		code := gjson.Get(lis.String(), "code").String()
		burl := gjson.Get(lis.String(), "url").String()
		rep, _ := http.Get(burl)
		json, _ := ioutil.ReadAll(rep.Body)
		list := gjson.Get(string(json), "list")
		for _, lit := range list.Array() {
			name := gjson.Get(lit.String(), "name").String()
			url := gjson.Get(lit.String(), "url").String()
			category := gjson.Get(lit.String(), "category").String()
			rank := gjson.Get(lit.String(), "rank").String()
			auth := gjson.Get(lit.String(), "auth").String()
			version := gjson.Get(lit.String(), "version").String()
			allinfo := "|" + name + "|" + url + "|" + category + "|" + rank + "|" + auth + "|" + version + "|" + author + "|" + code + "|"
			file = file + "\n" + allinfo
			println(allinfo)
			nums++
		}
	}
	file = "### 厚墨公共书源集中索引\n" + "#### 更新时间\n" + time + "\n" + "#### 已收录\n" + strconv.Itoa(nums) + "\n\n" + file
	os.MkdirAll("./out", os.ModePerm)
	ioutil.WriteFile("./out/README.md", []byte(file), 0640)
	println("已输出到 ./out/README.md")
}
