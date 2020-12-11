package main

import (
	"fmt"
	"github.com/dogukanayd/go-simple-pub-sub/internal/pub/pub"
	"github.com/dogukanayd/go-simple-pub-sub/internal/pub/sub"
	"github.com/google/uuid"
)

func main() {
	numberOfSubscribers := 3

	for i := 0; i < numberOfSubscribers; i++ {
		fmt.Printf("#%v ", i)
		go sub.Subscriber()
	}

	pub.Publish(uuid.New().String())
}
