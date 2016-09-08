package events

import (
	"errors"
	"github.com/codeship/go-retro"
	"log"
)

var ErrNetwork = retro.NewStaticRetryableError(errors.New("error: failed to connect"), 5, 10)

func CheckUnprocessedOrders() {
	err := retro.DoWithRetry(func() error {
		return checkRequest()
	})

	if err != nil {
		log.Fatal("Failed to send check.unprocessed.orders request %s\n", err.Error())
		return
	}

	log.Print("check.unprocessed.orders request sent")
}

func checkRequest() error {
	return nil
}
