package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var clusterClient *redis.ClusterClient

func main() {

	wait := sync.WaitGroup{}

	wait.Add(10)

	for i := 0; i < 10; i++ {
		go func() {

			// 连接redis集群
			client := redis.NewClusterClient(&redis.ClusterOptions{
				Addrs: []string{ // 填写master主机
					"172.16.1.2:7000",
					"172.16.1.2:7001",
					"172.16.1.2:7002",
				},
				ReadOnly:      true,
				RouteRandomly: true,
				Password:      "",               // 设置密码
				DialTimeout:   50 * time.Second, // 设置连接超时
				ReadTimeout:   50 * time.Second, // 设置读取超时
				WriteTimeout:  50 * time.Second, // 设置写入超时
			})

			client.Pipeline()
			// 发送一个ping命令,测试是否通
			//pong, err := client.Ping().Result()
			//fmt.Println(pong, err)
			s := time.Now()
			pipe := client.Pipeline()
			for j := 0; j < 10; j++ {

				//fmt.Println(client.Type("nearbypoi_wwfy8eh"))

				pipe.Get("1371648200")
				pipe.Get("1371648300")
				pipe.Get("1371648400")

				pipe.Get("nearbypoi_wwfy8eh")
				pipe.Get("nearbypoi_wt7ky80")
				pipe.Get("nearbypoi_wttw36k")
				pipe.Get("nearbypoi_wskm9vk")
				pipe.Get("nearbypoi_wkp18wd")
				pipe.Get("nearbypoi_wqdvf8r")
				pipe.Get("nearbypoi_wtu0cwh")
				pipe.Get("nearbypoi_ww468fx")
				pipe.Get("nearbypoi_wsdn8pn")
				pipe.Get("nearbypoi_ww0k57y")
				pipe.Get("nearbypoi_wt7ky80")

				cmders, _ := pipe.Exec()

				strArr := make([]string, 0, 20)

				for _, cmder := range cmders {
					cmd := cmder.(*redis.StringCmd)
					line, _ := cmd.Result()
					//if err != nil {
					//	fmt.Println("err", err)
					//}
					strArr = append(strArr, line)
				}

				//fmt.Println(strings.Join(strArr, " == "))
			}

			fmt.Println(time.Now().Sub(s).Seconds())

			wait.Done()
		}()
	}

	wait.Wait()

	time.Sleep(100000 * time.Second)

}
