package bloomfilter

import (
	"bufio"
	"fmt"
	"github.com/kevburnsjr/bloomfilter"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestRand1(t *testing.T) {

	num := 110000000

	//m, k := bloomfilter.EstimateParameters(num, 0.0000001)
	//bf := bloomfilter.New(m, k)
	//
	//for i := 0; i < num; i++ {
	//	str1 := "hello,bloom filter!" + strconv.Itoa(i)
	//	bf.Add([]byte(str1))
	//}
	//
	//fmt.Println(binary.Size(bf))
	//
	//fmt.Println(bf.Test([]byte("bar")))
	//
	//dst, _ := os.Create("/Users/byb/Downloads/bloom.data")
	//
	//dst.Write(bf.ToBytes())

	open, _ := os.Open("/Users/byb/Downloads/bloom.data")
	stat, _ := open.Stat()

	bt := make([]byte, stat.Size())

	open.Read(bt)
	fmt.Println(len(bt))
	open.Close()

	_, k1 := bloomfilter.EstimateParameters(num, 0.0000001)
	fmt.Println("dlajfakljfklsjflkjslkdj")
	fmt.Println(k1)
	fmt.Println("dlajfakljfklsjflkjslkdj")
	bf1 := bloomfilter.NewFromBytes(bt, k1)

	bt = nil

	fmt.Println(bf1.Test([]byte("92626952067511\tCMCC-XS3a")))
	fmt.Println(bf1.Test([]byte("220152333415537\t金泉宾馆5楼33")))

	fmt.Println("======================")

	var matex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			filepath := "/Users/byb/Downloads/bloom_test_data.txt"
			file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
			if err != nil {
				fmt.Println("Open file error!", err)
				return
			}

			buf := bufio.NewReader(file)

			s := time.Now()
			fmt.Println("s================")
			for {
				line, err := buf.ReadString('\n')
				line = strings.TrimSpace(line)
				matex.Lock()
				test := bf1.Test([]byte(line))
				matex.Unlock()
				if !test {
					fmt.Println(line)
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

			fmt.Println("e================")
			fmt.Println(time.Now().Sub(s).Seconds())
		}()
	}

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
