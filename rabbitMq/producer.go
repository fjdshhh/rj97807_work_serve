package main

import (
	"fmt"
	"rabbitMq/Tool"
	"strconv"
	"time"
)

func main() {
	rabbitmq := Tool.NewRabbitMqSubscription("list")
	for i := 0; i < 10000; i++ {
		rabbitmq.PublishSubscription("订阅模式生产者的第" + strconv.Itoa(i) + "条数据")
		fmt.Printf("订阅模式生产第" + strconv.Itoa(i) + "条数据\n")
		time.Sleep(1 * time.Second)
	}
}
