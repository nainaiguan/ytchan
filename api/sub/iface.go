package sub

import "context"

type Subchan interface {
	// Send send a message to the chan
	Send(message interface{}) error
	// Subscribe subscribe the chan and get the latest message when there's something new
	Subscribe(Name string, size int) chan interface{}
	// Unsubscribe unsubscribe the chan and you will not get the message from this chan
	Unsubscribe(Name string)
	// Size get the size of the chan
	Size() int
	// Capacity get the max size of the chan
	Capacity() int
	// History  get the latest history of the chan
	History() []interface{}
	// Close close the chan in a very safe and elegant way
	Close(cancelFunc context.CancelFunc)
}
