package main

/*
#include <windows.h>
#include <stdio.h>
#include <conio.h>

double FileTimeToDouble(FILETIME* pFiletime)
{
  return (double)((*pFiletime).dwHighDateTime * 4.294967296E9) + (double)(*pFiletime).dwLowDateTime;
}

double m_fOldCPUIdleTime;
double m_fOldCPUKernelTime;
double m_fOldCPUUserTime;

BOOL Initialize()
{
  FILETIME ftIdle, ftKernel, ftUser;
  BOOL flag = FALSE;
  if (flag = GetSystemTimes(&ftIdle, &ftKernel, &ftUser))
  {
	  m_fOldCPUIdleTime = FileTimeToDouble(&ftIdle);
	  m_fOldCPUKernelTime = FileTimeToDouble(&ftKernel);
	  m_fOldCPUUserTime = FileTimeToDouble(&ftUser);

  }
  return flag;
}

int GetCPUUseRate()
{
  int nCPUUseRate = -1;
  FILETIME ftIdle, ftKernel, ftUser;
  if (GetSystemTimes(&ftIdle, &ftKernel, &ftUser))
  {
	  double fCPUIdleTime = FileTimeToDouble(&ftIdle);
	  double fCPUKernelTime = FileTimeToDouble(&ftKernel);
	  double fCPUUserTime = FileTimeToDouble(&ftUser);
	  nCPUUseRate= (int)(100.0 - (fCPUIdleTime - m_fOldCPUIdleTime) / (fCPUKernelTime - m_fOldCPUKernelTime + fCPUUserTime - m_fOldCPUUserTime)*100.0);
	  m_fOldCPUIdleTime = fCPUIdleTime;
	  m_fOldCPUKernelTime = fCPUKernelTime;
	  m_fOldCPUUserTime = fCPUUserTime;
  }
  return nCPUUseRate;
}
int cpu()
{
  if (!Initialize())
  {
	  getch();
	  return -1;
  }
  else
  {
	  Sleep(100);
	  return GetCPUUseRate();
  }
  return -1;
}

DWORD getWin_MemUsage(){
    MEMORYSTATUS ms;
    GlobalMemoryStatus(&ms);
    return ms.dwMemoryLoad;
}
*/
import "C"
import (
	"fmt"
	"net/http"

	"github.com/lflxp/Wings/register"
)

func Memory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, fmt.Sprintf("%d", C.getWin_MemUsage()))
}

func main() {

	register.WatchDog()

	// cpu := C.cpu()
	// fmt.Printf("\r%d%%", cpu)
	// s := C.getWin_MemUsage()
	// fmt.Printf("%d", s)
	http.HandleFunc("/", Memory)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}

// // #define WIN32_LEAN_AND_MEAN
// // #include <windows.h>
// import "C"
// import "syscall"

// func GetCurrentDirectory() string {
// 	if bufLen := C.GetCurrentDirectoryW(0, nil); bufLen != 0 {
// 		buf := make([]uint16, bufLen)
// 		if bufLen := C.GetCurrentDirectoryW(bufLen, (*C.WCHAR)(&buf[0])); bufLen != 0 {
// 			return syscall.UTF16ToString(buf)
// 		}
// 	}
// 	return ""
// }

// func main() {
// 	println(GetCurrentDirectory())
// 	var a, b, c *C.struct__FILETIME
// 	println(C.GetSystemTimes(a, b, c))
// 	println(a, b, c)
// }

// package main

// import (
// 	"github.com/lflxp/Wings/register"
// 	// "github.com/lflxp/Wings/monitor/os"
// )

// func main() {
// 	wait := make(chan int)
// 	register.WatchDog()
// 	//启动监控api

// 	//启动rpc服务

// 	<-wait
// 	// fmt.Println(os.GetHostname())
// 	// fmt.Println(runtime.GOARCH)
// 	// fmt.Println(runtime.GOOS)
// 	// fmt.Println(fmt.Sprintf("%d",runtime.NumCPU))
// }
