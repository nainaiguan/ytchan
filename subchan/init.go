package subchan

import (
	"context"
	"ytChan/api/sub"
)

func Default() (*subChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	m := make(map[string]chan interface{})
	ret := &subChan{
		data:           make(chan interface{}, 1024),
		cap:            1024,
		maxSendProcess: 1024,
		subscriber: subscriber{
			m: m,
		},
		sendHistory: history{
			h: make([]interface{}, 0),
		},
		sendProcess: sendProcess{
			num: 0,
		},
		closeFlag: closeFlag{
			flag: 0,
		},
		ctx: ctx,
	}

	go ret.subChanCleanDaemon()

	return ret, shut
}

func New(args sub.NewSubArgs) (*subChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	m := make(map[string]chan interface{})
	ret := &subChan{
		data:           make(chan interface{}, args.Size),
		cap:            args.Size,
		maxSendProcess: args.MaxSendProcess,
		subscriber: subscriber{
			m: m,
		},
		sendHistory: history{
			h: make([]interface{}, 0),
		},
		sendProcess: sendProcess{
			num: 0,
		},
		closeFlag: closeFlag{
			flag: 0,
		},
		ctx: ctx,
	}

	go ret.subChanCleanDaemon()

	return ret, shut
}
