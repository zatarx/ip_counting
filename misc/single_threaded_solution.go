package misc

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var octets = [4]uint32{}

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

		ipMap[byteIndex] |= uint8(math.Pow(2.0, float64(bitIndex)))

	}

	uniqueIpCount := uint32(0)
	for _, byte_ := range ipMap {
		uniqueIpCount += uint32(bits.OnesCount8(byte_))
	}
	println(fmt.Sprintf("Total ip count %d", uniqueIpCount))

}
