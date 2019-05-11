package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	fmt.Println("您的系统:" + runtime.GOOS)                                    //windows
	fmt.Println(exec.Command("explorer", "http://localhost:4000").Start()) //nil
	files, err := ioutil.ReadDir("../")
	arrFileName := []string{}
	strName := ""
	if err != nil {
		fmt.Println("读取错误:", err)
		files, _ = ioutil.ReadDir("C:/")
	}
	for _, item := range files {
		arrFileName = append(arrFileName, item.Name())
		strName += item.Name() + ",: modeTime" +
			item.ModTime().Format("2006-01-02 15:04:05") + ",Mode:" + item.Mode().String() + ",大小:" +
			strconv.FormatInt(item.Size()/1024, 10) + "</br>"
	}
	//fmt.Println("item.Mode():" + arrFileName.Mode().Format("2006-01-02 15:04:05"))
	fmt.Println((files[0].Name()))
	fmt.Println(files[0].IsDir())
	fmt.Println(files[0].ModTime())
	//slice1 := arrFileName[0:]

	fmt.Println(strName)
	fmt.Printf("文件数量:%d", len(files))
	//strarrFileName := arrFileName[0]
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(strName))
		r.ParseForm()
		fmt.Printf("%v", r.Form["b"])
	})
	http.ListenAndServe(":4000", nil)
	// 测试git可视化工具
}
