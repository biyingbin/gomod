package bloomfilter

import (
	"bufio"
	"fmt"
	"github.com/kevburnsjr/bloomfilter"
	"io"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestRand1(t *testing.T) {

	num := 110000000

	m, k := bloomfilter.EstimateParameters(num, 0.0000001)
	bf := bloomfilter.New(m, k)

	dst, _ := os.Create("/Users/byb/Downloads/bloom.data")

	dst.Write(bf.ToBytes())

	fileInfoList, _ := ioutil.ReadDir("/home/datacenter/tmp/wplocserver")

	for f := range fileInfoList {
		filepath := "/home/datacenter/tmp/wplocserver/" + fileInfoList[f].Name()
		file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
		if err != nil {
			fmt.Println("Open file error!", err)
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}

		var size = stat.Size()
		fmt.Println("file size=", size)

		buf := bufio.NewReader(file)
		for {
			line, err := buf.ReadString('\n')
			//line = strings.TrimSpace(line)

			for i := 0; i < num; i++ {
				bf.Add([]byte(line))
			}

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

	}

	//open, _ := os.Open("/Users/byb/Downloads/bloom.data")
	//stat, _ := open.Stat()
	//
	//bt := make([]byte, stat.Size())
	//
	//open.Read(bt)
	//fmt.Println(len(bt))
	//open.Close()
	//
	//_, k1 := bloomfilter.EstimateParameters(num, 0.0000001)
	//bf1 := bloomfilter.NewFromBytes(bt, k1)
	//
	//bt = nil
	//
	//fmt.Println(bf1.Test([]byte("bar")))
	//fmt.Println(bf1.Test([]byte("sbar")))
	//
	//for i := 0; i < num; i++ {
	//	str1 := "hello,bloom filter!" + strconv.Itoa(i)
	//	test := bf1.Test([]byte(str1))
	//	if(!test) {
	//		fmt.Println("=============")
	//	}
	//}

	//fmt.Println("=======")
	//
	//fmt.Println(bf.Test([]byte("bar")))
	//fmt.Println(bf.Test([]byte("sbar")))
	//
	//fmt.Println("=======")
	//
	//s := time.Now()
	//
	//for i := 0; i < 10000 * 10000 + 1000000; i++ {
	//	str1 := "hello,bloom filter!" + strconv.Itoa(i)
	//	test := bf.Test([]byte(str1))
	//	_ = test
	//}
	//
	//fmt.Println(time.Now().Sub(s).Seconds())
	//fmt.Println("=======")
	//
	time.Sleep(time.Second * 100000)

}

//
//func TestMemoryBloomFilter(t *testing.T) {
//	var filter BloomFilter = NewMemoryBloomFilter(64<<20, 5)
//	RandTest(t, filter, 50000)
//
//}
//
//func TestFileBloomFilter(t *testing.T) {
//	target := "bloom.tmp"
//	defer os.Remove(target)
//	var filter BloomFilter = NewFileBloomFilter(target, 64<<20, 5)
//	RandTest(t, filter, 50000)
//}
