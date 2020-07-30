package main

//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//
//func main() {
//
//	wait := sync.WaitGroup{}
//
//	wait.Add(10)
//
//	for i := 0; i < 10; i++ {
//		go func() {
//
//			// 连接redis集群
//			client, err := redis.NewCluster(
//				&redis.Options{
//					StartNodes: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
//					ConnTimeout: 50 * time.Millisecond,
//					ReadTimeout: 50 * time.Millisecond,
//					WriteTimeout: 50 * time.Millisecond,
//					KeepAlive: 16,
//					AliveTime: 60 * time.Second,
//				})
//
//
//			client.Pipeline()
//			// 发送一个ping命令,测试是否通
//			//pong, err := client.Ping().Result()
//			//fmt.Println(pong, err)
//			s := time.Now()
//			pipe := client.Pipeline()
//			for j := 0; j < 10; j++ {
//
//				//fmt.Println(client.Type("nearbypoi_wwfy8eh"))
//
//				pipe.Get("1371648200")
//				pipe.Get("1371648300")
//				pipe.Get("1371648400")
//
//				pipe.Get("nearbypoi_wwfy8eh")
//				pipe.Get("nearbypoi_wt7ky80")
//				pipe.Get("nearbypoi_wttw36k")
//				pipe.Get("nearbypoi_wskm9vk")
//				pipe.Get("nearbypoi_wkp18wd")
//				pipe.Get("nearbypoi_wqdvf8r")
//				pipe.Get("nearbypoi_wtu0cwh")
//				pipe.Get("nearbypoi_ww468fx")
//				pipe.Get("nearbypoi_wsdn8pn")
//				pipe.Get("nearbypoi_ww0k57y")
//				pipe.Get("nearbypoi_wt7ky80")
//
//				cmders, _ := pipe.Exec()
//
//				strArr := make([]string, 0, 20)
//
//				for _, cmder := range cmders {
//					cmd := cmder.(*redis.StringCmd)
//					line, _ := cmd.Result()
//					//if err != nil {
//					//	fmt.Println("err", err)
//					//}
//					strArr = append(strArr, line)
//				}
//
//				//fmt.Println(strings.Join(strArr, " == "))
//			}
//
//			fmt.Println(time.Now().Sub(s).Seconds())
//
//			wait.Done()
//		}()
//	}
//
//	wait.Wait()
//
//	time.Sleep(100000 * time.Second)
//
//}
