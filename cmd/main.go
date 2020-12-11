package main

import (
	"github.com/dogukanayd/go-simple-pub-sub/internal/pub/pub"
	"github.com/dogukanayd/go-simple-pub-sub/internal/pub/sub"
	"github.com/google/uuid"
)

func main() {
	go func() {
		for {
			sub.Subscriber()
		}
	}()

	pub.Publish(uuid.New().String(), pub.Queue)
}
