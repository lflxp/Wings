package main

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
	n, _ := net.IOCounters(true)
	fmt.Println(n)

	d, _ := disk.Usage("c:/")
	fmt.Println(d)

	da, _ := disk.Partitions(true)
	fmt.Println(da)

	h, _ := host.Info()
	fmt.Println(h)
	fmt.Println(host.PlatformInformation())

	a, _ := load.Avg()
	fmt.Println(a)

	m, _ := load.Misc()
	fmt.Println(m)
}
