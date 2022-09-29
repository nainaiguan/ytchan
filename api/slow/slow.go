package slow

import "time"

type NewSlowArgs struct {
	Size           int
	Step           time.Duration
	MaxSendProcess int
}
