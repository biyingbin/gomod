package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
)

var clusterClient *redis.ClusterClient

func main() {

	fmt.Println("==================")

	wait := sync.WaitGroup{}

	wait.Add(50)

	for i := 0; i < 300; i++ {

		go func() {

			// 连接redis集群
			//client := redis.NewClusterClient(&redis.ClusterOptions{
			//	Addrs: []string{ // 填写master主机
			//		//"172.16.1.2:7000",
			//		"172.16.1.2:6379",
			//		//"172.16.1.2:7002",
			//	},
			//	Password:     "",               // 设置密码
			//	DialTimeout:  50 * time.Second, // 设置连接超时
			//	ReadTimeout:  50 * time.Second, // 设置读取超时
			//	WriteTimeout: 50 * time.Second, // 设置写入超时
			//})

			// 发送一个ping命令,测试是否通
			client := redis.NewClient(&redis.Options{
				Addr:     "172.16.1.2:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			})

			for {

				get := client.MGet("fjklajdfljfljaf")

				_ = get //.Result()

				//if (err != nil) {
				//	fmt.Println(err)
				//}

				//fmt.Println(get.Result())
				//fmt.Println("=======")
			}

		}()

	}

	wait.Wait()

}
