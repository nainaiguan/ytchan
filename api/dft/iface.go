package dft

import "context"

type Dftchan interface {
	// Send send a message to the chan
	Send(message interface{}) error
	// Pull pull a message from the top of the chan
	Pull(size int) []interface{}
	// Size get the size of the chan
	Size() int
	// Capacity get the max size of the chan
	Capacity() int
	// History  get the latest history of the chan
	History() []interface{}
	// Close close the chan in a very safe and elegant way
	Close(cancelFunc context.CancelFunc)
}
