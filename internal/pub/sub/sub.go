package sub

import (
	"github.com/dogukanayd/go-simple-pub-sub/internal/pub/pub"
	"github.com/dogukanayd/go-simple-pub-sub/internal/pub/req"
	"time"
)

// Subscriber subscribe the pub.Queue
func Subscriber() {
	for {
		val := <-pub.Queue
		requester := req.Requester{ID: val.(string)}
		requester.Send()

		time.Sleep(2 * time.Second)
	}
}
