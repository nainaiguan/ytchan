package dftchan

import "context"

func Default() (*DftChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	ret := &DftChan{
		data:           make(chan interface{}, 1024),
		cap:            1024,
		maxSendProcess: 1024,
		sendHistory: history{
			h: make([]interface{}, 0),
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
	go ret.DftChanCleanDaemon()
	return ret, shut
}
