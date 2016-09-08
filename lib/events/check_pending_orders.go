package events

import (
	"github.com/rafaeljesus/cron/lib/queue_manager"
	"github.com/streadway/amqp"
	"log"
)

func CheckPendingOrders() {
	chann := queue_manager.Channel
	err := chann.Publish("orders", "check_pending", true, false, amqp.Publishing{
		ContentType: "text/plain",
	})

	if err != nil {
		log.Fatalf("%s: %s", "failed to publish check.pending.orders message", err)
	}

	log.Print("check.pending.orders event message sent")
}
