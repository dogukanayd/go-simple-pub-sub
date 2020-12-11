package pub

import "time"

const (
	queueSize = 10
)

var Queue = make(chan interface{}, queueSize)

// Publish the data to the Queue channel
func Publish(message interface{}, ch chan interface{}) {
	ch <- message
	time.Sleep(time.Second * 1)
}
