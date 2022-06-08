package rmq

import (
	"fmt"
	"github.com/007team/douyin-micro/mq-server/settings"
	"github.com/streadway/amqp"
)

var MQ *amqp.Connection

func RabbitMQ(connString string) {
	conn,err := amqp.Dial(connString)
	if err != nil {
		panic(err)
	}
	MQ = conn
}

func Init(cfg *settings.RabbitmqConfig) (err error) {
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	cfg.User,
	//	cfg.Password,
	//	cfg.Host,
	//	cfg.Port,
	//	cfg.DbName,
	//)
	//pathRabbitMQ := strings.Join([]string{cfg.RabbitMQ, "://" , RabbitMQUser, ":", RabbitMQPassWord, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")
	pathRabbitMQ := fmt.Sprintf("%s://%s:%s@%s:%d/",
		cfg.Rabbitmq,
		cfg.RabbitmqUser,
		cfg.RabbitmqPassword,
		cfg.RabbitmqHost,
		cfg.RabbitmqPort,
		)
	fmt.Println(pathRabbitMQ)
	RabbitMQ(pathRabbitMQ)

	return err
}
