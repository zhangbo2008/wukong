/*

没有比这个更简单的例子了。

*/

package main

import (
	"fmt"
	"github.com/huichen/wukong/engine"
	"github.com/huichen/wukong/types"
	"io/ioutil"
	"log"
)

var (
	// searcher是线程安全的
	searcher1 = engine.Engine{}
)

func main() {

	//ReadFile
	//从这里看出来,go的路径和java一样都是默认针对项目根目录的.
	output, err := ioutil.ReadFile("data/dictionary.txt")
	if err != nil {
		fmt.Println("Read file error!")
		fmt.Println(err)
		return
	}
	//fmt.Println(string(output))











	// 初始化
	searcher1.Init(types.EngineInitOptions{
		SegmenterDictionaries: "data/dictionary.txt"})
	defer searcher1.Close()

	// 将文档加入索引，docId 从1开始
	searcher1.IndexDocument(1, types.DocumentIndexData{Content: "此次百度收购将成中国互联网最大并购"}, false)
	searcher1.IndexDocument(2, types.DocumentIndexData{Content: "百度宣布拟全资收购91无线业务"}, false)
	searcher1.IndexDocument(3, types.DocumentIndexData{Content: "百度是中国最大的搜索引擎"}, false)

	// 等待索引刷新完毕
	searcher1.FlushIndex()

	// 搜索输出格式见types.SearchResponse结构体
	log.Print(searcher1.Search(types.SearchRequest{Text: "百度中国"}))
}
