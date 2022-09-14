package slow

import "context"

type Slowchan interface {
	// Send send a message to the chan
	Send(message interface{}) error
	// Start start to pop out message
	Start()
	// Stop stop popping message
	Stop()
	// Size get the size of the chan
	Size() int
	// Capacity get the max size of the chan
	Capacity() int
	// History  get the latest history of the chan
	History() []interface{}
	// Close close the chan in a very safe and elegant way
	Close(cancelFunc context.CancelFunc)
}
