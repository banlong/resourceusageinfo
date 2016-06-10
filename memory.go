package main

//https://godoc.org/github.com/shirou/gopsutil has many ways modules to provide system information

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"resourceusageinfo/render"
)

func main() {
	v, _ := mem.VirtualMemory()
	fmt.Println("VIRTUAL MEMORY")
	fmt.Println("Total Memory: ", render.RenderInt64("#,###.",v.Total), "bytes")
	fmt.Println("Used        : ", render.RenderInt64("#,###.",v.Used), "bytes")
	fmt.Println("Used Percentage: : ", v.UsedPercent, "%")

	fmt.Println("CPU")
	cpus, _ := cpu.Counts(false)
	percet, _ := cpu.Percent(0, true)
	fmt.Println("CPUs: ", cpus)
	fmt.Println("Usage: ", percet, "%")

	timeStat, _ := cpu.Times(false)
	fmt.Println("CPU 0: ", timeStat[0])


}
