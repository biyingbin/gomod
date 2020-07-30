package file

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func TestFile(t *testing.T) {

	filepath := "/Users/byb/Downloads/polygon2.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}

	buf := bufio.NewReader(file)

	arr := make([][][][]float64, 1040397)

	fmt.Println("s================")

	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		var data [][][]float64

		json.Unmarshal([]byte(line), &data)

		arr = append(arr, data)

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

	time.Sleep(time.Second * 10000)
	fmt.Println("size = ", unsafe.Sizeof(&arr))
	time.Sleep(time.Second * 10000)

}
