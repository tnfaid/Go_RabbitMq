package main

import (
	"github.com/creamdog/gonfig"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	consumeRabbit()
}

func consumeRabbit() {
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

	conn, err := amqp.Dial(queueUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,    // name
		queueDurable, // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
