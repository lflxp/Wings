package main

import (
	"fmt"
	// "runtime"
	"github.com/lflxp/Wings/register"
	// "github.com/lflxp/Wings/monitor/os"
)

func main() {
	err := register.Register("conf/app.conf")
	if err != nil {
		fmt.Println("Error",err.Error())
		return
	}
	fmt.Println("ok")

	// fmt.Println(os.GetHostname())
	// fmt.Println(runtime.GOARCH)
	// fmt.Println(runtime.GOOS)
	// fmt.Println(fmt.Sprintf("%d",runtime.NumCPU))
}