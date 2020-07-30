package main

import (
	"bufio"
	"fmt"
	"github.com/kevburnsjr/bloomfilter"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main1() {

	num := 110000000

	m, k := bloomfilter.EstimateParameters(num, 0.0000001)
	bf := bloomfilter.New(m, k)

	funcName("/home/datacenter/tmp/wplocserver/8000_wifi_bssid_ssid", num, bf)
	funcName("/home/datacenter/tmp/wplocserver/wifi_bssid_ssid_bind_long", num, bf)

	dst, _ := os.Create("/home/datacenter/tmp/wplocserver/bloom.data")

	dst.Write(bf.ToBytes())

	dst.Close()

}

func funcName(dir string, num int, bf *bloomfilter.BloomFilter) {
	fileInfoList, _ := ioutil.ReadDir(dir)

	for f := range fileInfoList {
		filepath := dir + "/" + fileInfoList[f].Name()
		file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("Open file error!", err)
			return
		}

		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}

		var size = stat.Size()
		fmt.Println("file size=", size)

		buf := bufio.NewReader(file)
		for {
			line, err := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			bf.Add([]byte(line))

			if err != nil {
				if err == io.EOF {
					fmt.Println("File read ok!")
					break
				} else {
					fmt.Println("Read file error!", err)
					return
				}
			}
		}

		file.Close()
	}
}
