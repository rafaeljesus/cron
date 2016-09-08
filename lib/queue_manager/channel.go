package queue_manager

import (
	"github.com/streadway/amqp"
	"log"
)

var Channel *amqp.Channel

func Connect() error {
	c := make(chan *amqp.Error)

	go func() {
		err := <-c
		log.Println("reconnect: " + err.Error())
		Connect()
	}()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		panic("failed to connect message broker")
	}

	ch, err := conn.Channel()
	if err != nil {
		panic("failed to open a message broker channel")
	}

	conn.NotifyClose(c)
	ch.NotifyClose(c)

	declareExchange("orders", ch)
	declareExchange("downloads", ch)

	Channel = ch

	return nil
}

func declareExchange(name string, ch *amqp.Channel) {
	if err := ch.ExchangeDeclare(
		name,    // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	); err != nil {
		log.Fatalf("%s: %s", "failed to declare message broker exchange", err)
	}
}
