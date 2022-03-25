/**
 * @Author : admin
 * @File : main
 * @Date: 2022/3/23 22:20
 * @Description:
 */
package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println("os temp:", os.TempDir())
	fmt.Println("GetAbsPath:", GetCallerAbsPath())
}

func GetPathOrigin() string {
	// 获取当前项目运行绝对路径
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// 会将所有路径的符号链接都解析出来。
	//除此之外，它返回的路径，是直接可访问的。
	//如果 path 或返回值是相对路径，则是相对于进程当前工作目录。
	res, _ := filepath.EvalSymlinks(exePath)
	//fmt.Println("path:", res)
	return res
}

func GetCallerAbsPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("runtime.Caller error")
	}
	absPath := path.Dir(filename)
	return absPath
}
