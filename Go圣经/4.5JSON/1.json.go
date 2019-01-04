package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string // 使用`...`进行标签定义
	Year   int    `json:"released"`        // 在JSON中以released替换Year
	Color  bool   `json:"color,omitempty"` // 当此项为true的时候才显示出来
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey", "Ingrid"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: false,
		Actors: []string{"Paul"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve", "Jacqueline"}}}

func main() {
	//// 将Go的数据结构转化为JSON
	//data, err := json.Marshal(movies)
	//if err != nil { log.Fatalf("JSON marshaling failed: %s", err) }

	// 上面的方法不容易阅读 因为把所有的信息都浓缩在了一个slice里
	// 所以利用以下方法更易于阅读
	NewData, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", NewData)

	// 通过unmarshal对JSON字符串进行解码
	// 下面将JSON数据结构转换到结构体slice中 这个结构体唯一的成员就是Title
	// 通过合理的定义Go中的数据结构 我们可以选择将哪部分数据解码到结构体中
	var titles []struct{ Title string }
	if err := json.Unmarshal(NewData, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)

	// PS:
	// json包中还有流式编码解码器 json.Encoder json.Decoder
}
