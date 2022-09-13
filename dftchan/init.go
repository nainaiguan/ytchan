package dftchan

import (
	"context"
	"ytChan/api/dft"
)

func Default() (*DftChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	ret := &DftChan{
		data:           make(chan interface{}, 1024),
		cap:            1024,
		maxSendProcess: 1024,
		sendHistory: history{
			max: 1024,
			h:   make([]interface{}, 0),
		},
		sendProcess: sendProcess{
			num: 0,
		},
		pullProcess: pullProcess{
			num: 0,
		},
		closeFlag: closeFlag{
			flag: 0,
		},
		ctx: ctx,
	}
	go ret.dftChanCleanDaemon()
	return ret, shut
}

func New(args dft.NewDftArgs) (*DftChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	ret := &DftChan{
		data:           make(chan interface{}, args.Size),
		cap:            args.Size,
		maxSendProcess: args.MaxSendProcess,
		sendHistory: history{
			max: args.MaxHistory,
			h:   make([]interface{}, 0),
		},
		sendProcess: sendProcess{
			num: 0,
		},
		pullProcess: pullProcess{
			num: 0,
		},
		closeFlag: closeFlag{
			flag: 0,
		},
		ctx: ctx,
	}
	go ret.dftChanCleanDaemon()
	return ret, shut
}
