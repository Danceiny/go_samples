package main

import (
	"fmt"
	"github.com/buger/jsonparser"
)
func main()  {

	data := []byte(`{
	  "person": {
		"name": {
		  "first": "Leonid",
		  "last": "Bugaev",
		  "fullName": "Leonid Bugaev"
		},
		"github": {
		  "handle": "buger",
		  "followers": 109
		},
		"avatars": [
		  { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
		]
	  },
	  "company": {
		"name": "Acme"
	  }
	}`)

	a := Pic(10, 20)
	fmt.Print(a)
	// You can specify key path by providing arguments to Get function

	value, dataType, offSet, err := jsonparser.Get(data, "person", "name", "fullName")
	if err != nil{
		fmt.Println("error")
	}
	var fullName string;
	if dataType == jsonparser.String {
		fullName = string(value)

	}
	fmt.Println(offSet)
	fmt.Println(fullName)


	paths := [][]string{
		[]string{"person", "name", "fullName"},
		[]string{"person", "avatars", "[0]", "url"},
		[]string{"company", "url"},
	}
	jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, err error){
		switch idx {
		case 0: // []string{"person", "name", "fullName"}
			println('0')
			fmt.Println('0')
			v, e := jsonparser.ParseString(value)
			if e == nil{
				println(v)
			}
		case 1: // []string{"person", "avatars", "[0]", "url"}
			println('1')
			fmt.Println('1')
		case 2: // []string{"company", "url"},
			println('2')
			fmt.Println('2')
		}
	}, paths...)



}
//func main() {

	//// 创建一个硬链接。
	//// 创建后同一个文件内容会有两个文件名，改变一个文件的内容会影响另一个。
	//// 删除和重命名不会影响另一个。
	//err := os.Link("main.go", "original_also.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("creating sym")
	//
	//// Create a symlink
	//err = os.Symlink("main.go", "original_sym.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// Lstat返回一个文件的信息，但是当文件是一个软链接时，它返回软链接的信息，而不是引用的文件的信息。
	//// Symlink在Windows中不工作。
	//fileInfo, err := os.Lstat("original_sym.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("Link info: %+v", fileInfo)
	////改变软链接的拥有者不会影响原始文件。
	//err = os.Lchown("original_sym.txt", os.Getuid(), os.Getgid())
	//if err != nil {
	//	log.Fatal(err)
	//}
//}
