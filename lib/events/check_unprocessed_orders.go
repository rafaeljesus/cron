package events

import (
	"github.com/rafaeljesus/cron/lib/queue_manager"
	"github.com/streadway/amqp"
	"log"
)

func CheckUnprocessedOrders() {
	chann := queue_manager.Channel
	err := chann.Publish("orders", "check_unprocessed", true, false, amqp.Publishing{
		ContentType: "text/plain",
	})

	if err != nil {
		log.Fatalf("%s: %s", "failed to publish check.unprocessed.orders message", err)
	}

	log.Print("check.unprocessed.orders event message sent")
}
