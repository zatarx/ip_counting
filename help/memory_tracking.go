package help

import (
	"fmt"
	"runtime"
	"math"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func translateIpToDecimal(octets [4]uint32) uint32 {
	return octets[3] + octets[2]*256 + octets[1]*(uint32(math.Pow(256.0, 2))) + octets[0]*uint32(math.Pow(256.0, 3))
}