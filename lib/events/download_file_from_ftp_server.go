package events

import (
	"github.com/rafaeljesus/cron/lib/queue_manager"
	"github.com/streadway/amqp"
	"log"
)

func DownloadFileFromFtpServer() {
	chann := queue_manager.Channel
	err := chann.Publish("downloads", "from_ftp_server", true, false, amqp.Publishing{
		ContentType: "text/plain",
	})

	if err != nil {
		log.Fatalf("%s: %s", "failed to publish download.file.from.server message", err)
		return
	}

	log.Print("download.file.from.server request sent")
}
