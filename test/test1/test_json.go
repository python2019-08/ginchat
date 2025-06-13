package test1

import (
	"encoding/json"
	"fmt"
	"os"
)

//	type Website struct {
//		Name   string `xml:"name,attr"`
//		Url    string
//		Course []string
//	}
type Website struct {
	Name   string   `json:"nameXX"`
	Url    string   `json:"urlXX"`
	Course []string `json:"courseXX"`
}

func Demo_writeJson() {
	info := []Website{
		{"Golang",
			"http://c.biancheng.net/golang/",
			[]string{"http://c.biancheng.net/cplus/",
				"http://c.biancheng.net/linux_tutorial/"}},
		{"Java",
			"http://c.biancheng.net/java/",
			[]string{"http://c.biancheng.net/socket/",
				"http://c.biancheng.net/python/",
				"http://c11.biancheng.net/python/",
				"http://c12.biancheng.net/python/"}},
		{"c#",
			"http://csharp.biancheng.net/java/",
			[]string{"http://csharp1.biancheng.net/socket/",
				"http://csharp2.biancheng.net/python/",
				"http://csharp3.biancheng.net/python/",
				"http://csharp4.biancheng.net/python/"}}}

	// 创建文件
	filePtr, err := os.Create("test1/info.json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()

	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)

	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误", err.Error())

	} else {
		fmt.Println("编码成功")
	}
}

func Demo_readJson() {
	filePtr, err := os.Open("./info.json")
	if err != nil {
		fmt.Println("文件打开失败 [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	var info []Website
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}
