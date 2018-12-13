package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	// 出于安全考虑，在 web 应用中最好使用json.MarshalforHTML()函数，
	// 其对数据执行HTML转码，所以文本可以被安全地嵌在HTML <script> 标签中
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)

	// 要想把JSON直接写入文件, 可以使用json.NewEncoder初始化文件(或者任何实现io.Writer 的类型), 并调用
	// Encode(); 反过来与其对应的是使用json.Decoder 和Decode()函数：
	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file) // NewEncoder returns a new encoder that writes to w.
	err := enc.Encode(vc)        // Encode writes the JSON encoding of v to the stream
	if err != nil {
		log.Println("Error in encoding json")
	}
}
