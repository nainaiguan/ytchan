package slowchan

import (
	"context"
	"time"
	"ytChan/api/slow"
)

func Default() (*SlowChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	ret := &SlowChan{
		data:           make(chan interface{}, 1024),
		cap:            1024,
		step:           1 * time.Second,
		maxSendProcess: 1024,
		sendHistory: history{
			max: 1024,
			h:   make([]interface{}, 0),
		},
		sendProcess: sendProcess{
			num: 0,
		},
		closeFlag: closeFlag{
			flag: 0,
		},
		ctx: ctx,
	}
	go ret.slowChanCleanDaemon()

	return ret, shut
}

func New(args slow.NewSlowArgs) (*SlowChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	ret := &SlowChan{
		data:           make(chan interface{}, args.Size),
		cap:            args.Size,
		step:           args.Step,
		maxSendProcess: args.MaxSendProcess,
		sendHistory: history{
			max: args.MaxHistory,
			h:   make([]interface{}, 0),
		},
		sendProcess: sendProcess{
			num: 0,
		},
		closeFlag: closeFlag{
			flag: 0,
		},
		ctx: ctx,
	}
	go ret.slowChanCleanDaemon()

	return ret, shut
}
