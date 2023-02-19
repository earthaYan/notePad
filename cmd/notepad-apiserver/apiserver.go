package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/earthaYan/notePad/internal/apiserver"
)

func main() {
	// 为每个程序生成一个固定的随机值,避免计算机程序的伪随机
	rand.Seed(time.Now().UTC().UnixNano())
	// GOMAXPROCS——逻辑CPU数量
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		// 使用 runtime.NumCPU() 查询 CPU 数量，并使用 runtime.GOMAXPROCS() 函数进行设置
		// 让代码并发执行
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	// NewApp方法使用默认参数创一个App对象
	apiserver.NewApp("notepad-apiserver").Run()
}
