package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

func convertIpToDecimal(octets [4]uint32) uint32 {
	return (octets[3] + octets[2]*256 + octets[1]*(uint32(math.Pow(256.0, 2))) +
		octets[0]*uint32(math.Pow(256.0, 3)))
}

func processChunk(ipChunk []string, uniqueIpMap []uint8, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for _, strIp := range ipChunk {
		var ipOctets = [4]uint32{}

		strIpOctets := strings.Split(strIp, ".")

		for index, octetStr := range strIpOctets {
			octet, err := strconv.Atoi(octetStr)

			if octet > 255 {
				println(
					fmt.Sprintf("Fix the input file. "+
						"Octet shouldn't be greater than 255, but got %d", octet),
				)
			} else if err != nil {
				println(
					fmt.Sprintf(
						"Error occurred while converting an octet. "+
							"Error: %s\n Initial value: %s\n Converted value: %d\n Ip: %s",
						err.Error(), octetStr, octet, strIp,
					),
				)
			}

			ipOctets[index] = uint32(octet)
		}

		decimalIp := convertIpToDecimal(ipOctets)
		byteIndex, bitIndex := decimalIp/8, decimalIp%8

		mutex.Lock()
		uniqueIpMap[byteIndex] |= uint8(math.Pow(2.0, float64(bitIndex)))
		mutex.Unlock()
	}
}

func main() {
	var fileName string
	var wg sync.WaitGroup
	var mutex sync.Mutex
	const IPV4_BITMAP_SIZE, IP_CHUNK_SIZE, DEFAULT_FILENAME = 536870912, 1000, "ip_addresses"
	runtime.GOMAXPROCS(8)

	print("Enter a filename (default ip_addresses): ")
	_, err := fmt.Scanln(&fileName)
	if err != nil {
		println("Error while reading the file:", err.Error())
	}

	if strings.TrimSpace(fileName) == "" {
		fileName = DEFAULT_FILENAME
	}

	println("Using file name: ", fileName)

	file, err := os.Open(fileName)
	if err != nil {
		println(fmt.Sprintf("Error opening file: %s", err.Error()))
		return
	}
	defer file.Close()

	start := time.Now()

	uniqueIpMap := new([IPV4_BITMAP_SIZE]uint8)

	ipsChunk := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ipStr := scanner.Text()
		ipsChunk = append(ipsChunk, ipStr)

		if len(ipsChunk) == IP_CHUNK_SIZE {
			wg.Add(1)
			go processChunk(ipsChunk[:], uniqueIpMap[:], &mutex, &wg)
			ipsChunk = make([]string, 0)
		}
	}

	if len(ipsChunk) != 0 {
		wg.Add(1)
		go processChunk(ipsChunk, uniqueIpMap[:], &mutex, &wg)
	}
	wg.Wait()

	var uniqueIpCount uint32 = 0
	for _, byte_ := range uniqueIpMap {
		uniqueIpCount += uint32(bits.OnesCount8(byte_))
	}

	elapsed := time.Since(start).Seconds()
	println(fmt.Sprintf("Total ip count %d, time passed: %f seconds", uniqueIpCount, elapsed))

}
