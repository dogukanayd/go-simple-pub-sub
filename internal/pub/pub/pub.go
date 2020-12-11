package pub

const (
	queueSize = 10
)

var Queue = make(chan interface{}, queueSize)

// Publish the data to the Queue channel
func Publish(data interface{}) {
	for {
		Queue <- data
	}
}
