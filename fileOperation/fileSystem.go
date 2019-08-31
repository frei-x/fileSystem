package fileOperation

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

var StrName string

func init() {
	userInfo, _ := user.Current()
	fmt.Println(userInfo)
	fmt.Println(userInfo.Username + ",您的系统:" + runtime.GOOS) //windows
	files, err := ioutil.ReadDir("./")
	arrFileName := []string{}
	if err != nil {
		fmt.Println("读取错误:", err)
		files, _ = ioutil.ReadDir("C:/")
	}
	for _, item := range files {
		arrFileName = append(arrFileName, item.Name())
		StrName += item.Name() + ",: modeTime" +
			item.ModTime().Format("2006-01-02 15:04:05") + ",Mode:" + item.Mode().String() + ",大小:" +
			strconv.FormatInt(item.Size()/1024, 10) + "</br>"
	}
	//fmt.Println("item.Mode():" + arrFileName.Mode().Format("2006-01-02 15:04:05"))
	fmt.Println((files[0].Name()))
	fmt.Println(files[0].IsDir())
	fmt.Println(files[0].ModTime())
	fmt.Println("sys信息:")
	fmt.Println(files[1].Sys())
	fileInfo, _ := os.Stat("./main.exe")
	//Truncate:截断, 修改文件大小,如果目标值小于原文件 则文件可能损坏
	setFileSize := os.Truncate("./main.exe", 1) // byte *
	if setFileSize != nil {
		fmt.Println(setFileSize)
	}
	// 修改文件修改时间与访问时间 / unix 秒
	setFileTime := os.Chtimes("./README.md", time.Unix(897084366, 0), time.Unix(16597723666666, 0))
	fmt.Println(setFileTime)
	fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData) //win 系统调度
	fmt.Println(fileSys.LastWriteTime.Nanoseconds())            //  纳秒 /1e6 为毫秒 最后修改时间
	fmt.Println((fileSys.CreationTime.Nanoseconds()))           //创建时间
	//	fmt.Println(time.Parse("2006-01-02 03:04:05", fileSys.CreationTime.Nanoseconds()/1e9)
	fmt.Println(fileSys.LastAccessTime.Nanoseconds()) //最后访问时间
	//slice1 := arrFileName[0:]
	fmt.Printf("文件数量:%d", len(files))
	exec.Command("explorer", "http://localhost:4000").Start()
}
