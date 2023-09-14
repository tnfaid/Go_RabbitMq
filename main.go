package main

import (
	"github.com/creamdog/gonfig"
	rabbitmq "goland-tutorial-progress/RabbitMq"
	rediscli "goland-tutorial-progress/RedisCli"
	"os"
)

func main() {
	f, err := os.Open("myconfig.json")
	if err != nil {
		// TODO: error handling
	}
	defer f.Close()
	config, err := gonfig.FromJson(f)
	if err != nil {
		// TODO: error handling
	}
	queueName, err := config.GetString("rabbit/queue_name", "null")
	queueUrl, err := config.GetString("rabbit/queue_url", "null")
	queueDurable, err := config.GetBool("rabbit/queue_durable", false)
	redisUrl, err := config.GetString("redis/redis_url", false)

	//rabbitmq.PublishRabbit(queueName, queueUrl, queueDurable)
	rabbitmq.ConsumeRabbit(queueName, queueUrl, queueDurable)
	rediscli.RedisHash(redisUrl)
	//rediscli.RedisNormal()

}
