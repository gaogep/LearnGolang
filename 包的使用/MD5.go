package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("zpf"))
	fmt.Println(h)
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr))
}

/*

func main() {
str := "abc123"

//方法一
data := []byte(str)
has := md5.Sum(data)
md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

fmt.Println(md5str1)

//方法二

w := md5.New()
io.WriteString(w, str)   //将str写入到w中
md5str2: = fmt.Sprintf("%x", w.Sum(nil))  //w.Sum(nil)将w的hash转成[]byte格式

fmt.Println(mdtstr2)
}

*/
