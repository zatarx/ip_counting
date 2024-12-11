package misc

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	// "runtime"
	// "lightspeed/misc"
	"strconv"
	"strings"
	"sync"
	"time"
)

// use multiprocessing (goroutines)
// use map to get through all the files
// 4294967296 options total (each bit represents an option)
// 407235279
// aka 256**4

func translateIpToDecimalv2(octets [4]uint32) uint32 {
	return octets[3] + octets[2]*256 + octets[1]*(uint32(math.Pow(256.0, 2))) + octets[0]*uint32(math.Pow(256.0, 3))
}

func processChunk(ipChunk []string, ipMap []uint8, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, strIp := range ipChunk {
		var octets = [4]uint32{}

		strOctets := strings.Split(strIp, ".")

		for index, octetStr := range strOctets {
			octet, err := strconv.Atoi(octetStr)
			if octet > 255 {
				panic(fmt.Sprintf("Octet shouldn't be greater than 255, but got %d", octet))
			} else if err != nil {
				println(err.Error())
				println(octet)
				println(octetStr)
				println(strIp)
			}

			octets[index] = uint32(octet)
		}
		// println(strIp)

		decimalIp := translateIpToDecimalv2(octets)
		byteIndex, bitIndex := decimalIp/8, decimalIp%8

		// println("decimal ip byte index and bit index:", byteIndex, bitIndex, decimalIp)
		mutex.Lock()
		ipMap[byteIndex] |= uint8(math.Pow(2.0, float64(bitIndex)))
		mutex.Unlock()

		// PrintMemUsage()
		// runtime.GC()
	}
}

func concurrentMain() {
	start := time.Now()
	PrintMemUsage()
	// value, err := strconv.Atoi("0")
	// println(value)

	file, err := os.Open("4294967296_ips.txt")
	// file, err := os.Open("ip_file")
	if err != nil {
		fmt.Println("error", err)
	}
	defer file.Close()

	// 4294967296 / 8 (total amount of bytes needed to store ip addresses)
	uniqueIpMap := new([536870912]uint8)
	println("after ipmap declaration")

	PrintMemUsage()

	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	var mutex sync.Mutex

	ips := make([]string, 0)

	for scanner.Scan() {
		ipStr := scanner.Text()
		ips = append(ips, ipStr)

		if len(ips) == 1000 {
			wg.Add(1)
			go processChunk(ips[:], uniqueIpMap[:], &mutex, &wg)
			ips = make([]string, 0)
			// ips = ips[:0]
		}
	}

	if len(ips) != 0 {
		wg.Add(1)
		go processChunk(ips, uniqueIpMap[:], &mutex, &wg)
	}
	wg.Wait()

	uniqueIpCount := uint32(0)
	for _, byte_ := range uniqueIpMap {
		uniqueIpCount += uint32(bits.OnesCount8(byte_))
	}
	PrintMemUsage()
	elapsed := time.Since(start).Seconds()
	println(fmt.Sprintf("Total ip count %d, time passed: %f seconds", uniqueIpCount, elapsed))

}

func main() {
	// singleThreadedMain()
	// help.ArraysSlices()
	concurrentMain()
}
