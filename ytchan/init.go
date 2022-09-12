package ytchan

import (
	"context"
	"ytChan/api/dftchan"
)

func Default() (*dftchan.DftChan, context.CancelFunc) {
	ctx, shut := context.WithCancel(context.Background())
	ret := &dftchan.DftChan{
		Data: make(chan interface{}, 1024),
		Cap:  1024,
		History: dftchan.History{
			H: make([]interface{}, 1024),
		},
		SendProcess: dftchan.SendProcess{
			Num: 0,
		},
		PullProcess: dftchan.PullProcess{
			Num: 0,
		},
		CleanFlag: dftchan.CleanFlag{
			Flag: 0,
		},
		CloseFlag: dftchan.CloseFlag{
			Flag: 0,
		},
		Ctx: ctx,
	}
	go ret.DftChanCleanDaemon()
	return ret, shut
}
