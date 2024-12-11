package help

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"math/bits"
)

func singleThreadedMain() {
	PrintMemUsage()
	file, err := os.Open("ip_file")
	if err != nil {
		fmt.Println("error", err)
	}
	defer file.Close()

	PrintMemUsage()

	// 4294967296 / 8 (total amount of bytes needed to store ip addresses)
	ipMap := new([536870912]uint8)
	println("after ipmap declaration")

	PrintMemUsage()

	// uniqueIps := make(map[string]struct{}) // map with empty keys
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var octets = [4]uint32{}
		// sem := make(chan struct{}, 3)
		if len(octets) > 0 {

		}

		strOctets := strings.Split(scanner.Text(), ".")

		for index, octetStr := range strOctets {
			octet, err := strconv.Atoi(octetStr)
			if octet > 255 {
				panic(fmt.Sprintf("Octet shouldn't be greater than 255, but got %d", octet))
			} else if err != nil {
				panic(err)
			}

			octets[index] = uint32(octet)
		}

		decimalIp := translateIpToDecimal(octets)
		byteIndex, bitIndex := decimalIp/8, decimalIp%8

		// println("decimal ip byte index and bit index:", byteIndex, bitIndex, decimalIp)

		ipMap[byteIndex] |= uint8(math.Pow(2.0, float64(bitIndex)))

		// PrintMemUsage()
		runtime.GC()
		// println("byte index:", byteIndex, "byte value: ",
		// 	ipMap[byteIndex], "index new value: ", ipMap[byteIndex]|uint8(math.Pow(2.0, float64(bitIndex))))
		// ipAddr, err := strconv.Atoi(ipStr)

	}

	uniqueIpCount := uint32(0)
	for _, byte_ := range ipMap {
		uniqueIpCount += uint32(bits.OnesCount8(byte_))
	}
	println(fmt.Sprintf("Total ip count %d", uniqueIpCount))

}
