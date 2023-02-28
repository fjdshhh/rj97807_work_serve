package Tool

import (
	"embed"
	"fmt"
	"github.com/streadway/amqp"
	"gopkg.in/ini.v1"
)

var MQURL = ""

//go:embed "config.ini"
var config embed.FS

func init() {
	data, err := config.ReadFile("config.ini")
	file, err := ini.Load(data)
	if err != nil {
		fmt.Println("配置文件错误", err)
	}
	LoadData(file)
}
func LoadData(file *ini.File) {
	MQURL = file.Section("RabbitMq").Key("Url").MustString("")
}

// RabbitMq RabbitMq结构体
type RabbitMQ struct {
	//连接
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列
	QueueName string
	//交换机名称
	ExChange string
	//绑定的key名称
	Key string
	//连接的信息，上面已经定义好了
	MqUrl string
}

func NewRabbitMQ(queueName string, exChange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, ExChange: exChange, Key: key, MqUrl: MQURL}
}

//关闭conn和chanel的方法
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误的函数处理
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		fmt.Printf("err是:%s,发送的信息是:%s", err, message)
	}
}

// NewRabbitMqSubscription 获取mq实例 -- 订阅模式
func NewRabbitMqSubscription(exchangeName string) *RabbitMQ {
	//创建rabbitmq实例
	//exchangeName相当于建立的通道 消费者和生产者都要在这个通道里
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "订阅模式连接rabbitmq失败。")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "订阅模式获取channel失败")
	return rabbitmq
}

// PublishSubscription 订阅模式发布消息
func (r *RabbitMQ) PublishSubscription(message string) {
	//第一步，尝试连接交换机
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"fanout", //这里一定要设计为"fanout"也就是广播类型。
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式发布方法中尝试连接交换机失败。")
	//第二步，发送消息
	err = r.channel.Publish(
		r.ExChange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}
func (r *RabbitMQ) ConsumeSbuscription() {
	//第一步，试探性创建交换机exchange
	err := r.channel.ExchangeDeclare(
		r.ExChange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式消费方法中创建交换机失败。")

	//第二步，试探性创建队列queue
	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "订阅模式消费方法中创建创建队列失败。")

	//第三步，绑定队列到交换机中
	err = r.channel.QueueBind(
		q.Name,
		"", //在pub/sub模式下key要为空
		r.ExChange,
		false,
		nil,
	)

	//第四步，消费消息
	messages, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range messages {
			fmt.Printf("订阅模式下收到的消息：%s\n", d.Body)
		}
	}()

	fmt.Println("订阅模式消费者已开启")
	<-forever

}
