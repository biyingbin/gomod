package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var clusterClient *redis.ClusterClient

func t() {

	wait := sync.WaitGroup{}

	wait.Add(50)

	go func() {

		// 连接redis集群
		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{ // 填写master主机
				//"172.16.1.2:7000",
				"172.16.1.2:7001",
				//"172.16.1.2:7002",
			},
			Password:     "",               // 设置密码
			DialTimeout:  50 * time.Second, // 设置连接超时
			ReadTimeout:  50 * time.Second, // 设置读取超时
			WriteTimeout: 50 * time.Second, // 设置写入超时
		})
		// 发送一个ping命令,测试是否通
		pong, err := client.Ping().Result()
		fmt.Println(pong, err)

		for {
			get := client.Get("aaaa")

			_ = get
			//fmt.Println(get.Result())
			//fmt.Println("=======")
		}

	}()

	wait.Wait()

}
