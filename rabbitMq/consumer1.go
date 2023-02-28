package main

import "rabbitMq/Tool"

func main() {
	rabbitmq := Tool.NewRabbitMqSubscription("list")
	rabbitmq.ConsumeSbuscription()
}
