package subchan

import (
	"context"
	"ytChan/api/sub"
)

func Default() (*SubChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	m := make(map[string]chan interface{})
	ret := &SubChan{
		data:           make(chan interface{}, 1024),
		cap:            1024,
		maxSendProcess: 1024,
		subscriber: subscriber{
			m: m,
		},
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

	go ret.SubChanCleanDaemon()
	go ret.Reconcile()

	return ret, shut
}

func New(args sub.NewSubArgs) (*SubChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	m := make(map[string]chan interface{})
	ret := &SubChan{
		data:           make(chan interface{}, args.Size),
		cap:            args.Size,
		maxSendProcess: args.MaxSendProcess,
		subscriber: subscriber{
			m: m,
		},
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

	go ret.SubChanCleanDaemon()
	go ret.Reconcile()

	return ret, shut
}
