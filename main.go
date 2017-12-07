package main

import (
	"fmt"
	// "runtime"
	"github.com/lflxp/Wings/register"
	// "github.com/lflxp/Wings/monitor/os"
)

func main() {
	//注册服务
	err := register.Register("conf/app.conf")
	if err != nil {
		fmt.Println("Error",err.Error())
		return
	}
	fmt.Println("ok")
	//启动监控api

	//启动rpc服务

	// fmt.Println(os.GetHostname())
	// fmt.Println(runtime.GOARCH)
	// fmt.Println(runtime.GOOS)
	// fmt.Println(fmt.Sprintf("%d",runtime.NumCPU))
}