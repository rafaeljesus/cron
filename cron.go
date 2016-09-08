package main

import (
	"log"
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/rafaeljesus/cron/lib/events"
	"github.com/rafaeljesus/cron/lib/healthz"
	"github.com/rafaeljesus/cron/lib/queue_manager"
)

func main() {
	queue_manager.Connect()

	gocron.Every(1).Second().Do(events.CheckPendingOrders)
	gocron.Every(1).Second().Do(events.CheckUnprocessedOrders)
	gocron.Every(1).Day().Do(events.DownloadFileFromFtpServer)

	<-gocron.Start()

	http.HandleFunc("/healthz", healthz.Index)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
