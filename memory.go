package main

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
	"resourceusageinfo/render"
)

func main() {
	v, _ := mem.VirtualMemory()
	s, _ := mem.SwapMemory()
	// almost every return value is a struct
	fmt.Println("VIRTUAL MEMORY")
	fmt.Println("Total Memory: ", render.RenderInt64("#,###.",v.Total), "bytes")
	fmt.Println("Used        : ", render.RenderInt64("#,###.",v.Used), "bytes")
	fmt.Println("Used Percentage: : ", v.UsedPercent, "%")

	fmt.Println("SWAP MEMORY")
	fmt.Println("Total Memory: ", render.RenderInt64("#,###.",s.Total), "bytes")
	fmt.Println("Used        : ", render.RenderInt64("#,###.",s.Used), "bytes")
	fmt.Println("Used Percentage: : ", s.UsedPercent, "%")
}
