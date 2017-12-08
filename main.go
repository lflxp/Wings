package main

import (
	"fmt"
	// "runtime"
	"time"
	"github.com/lflxp/Wings/register"
	// "github.com/lflxp/Wings/monitor/os"
)

func Go() {
	err := register.Register("conf/app.conf")
	if err != nil {
		fmt.Println("Error",err.Error())
		return
	}
}

func WatchDog() {
	Go()
	gcInterval,_ := time.ParseDuration("29s")
	ticker := time.NewTicker(gcInterval)
	go func() {
		for {
			select {
			case <- ticker.C:
				Go()
			}
		}
	}()
}

func main() {
	wait := make(chan int)
	WatchDog()

	<-wait
	// fmt.Println(os.GetHostname())
	// fmt.Println(runtime.GOARCH)
	// fmt.Println(runtime.GOOS)
	// fmt.Println(fmt.Sprintf("%d",runtime.NumCPU))
}