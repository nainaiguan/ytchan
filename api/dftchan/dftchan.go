package dftchan

import (
	"context"
	"sync"
	"time"
)

type DftChan struct {
	Data        chan interface{}
	Cap         int
	History     History
	SendProcess SendProcess
	PullProcess PullProcess
	CleanFlag   CleanFlag
	CloseFlag   CloseFlag
	Ctx         context.Context
}

func (d *DftChan) DftChanCleanDaemon() {
	for {
		select {
		case <-d.Ctx.Done():
			return
		default:
			d.CleanFlag.Clean()
			l := len(d.History.H)
			tmp := make([]interface{}, l)
			copy(tmp, d.History.H)
			d.History.H = tmp
			d.CleanFlag.Done()
			time.Sleep(30 * time.Second)
		}
	}
}

type CleanFlag struct {
	Flag int
	M    sync.RWMutex
}

func (c *CleanFlag) Load() int {
	c.M.RLock()
	ret := c.Flag
	c.M.RUnlock()
	return ret
}

func (c *CleanFlag) Clean() {
	c.M.Lock()
	c.Flag = 1
	c.M.Unlock()
}

func (c *CleanFlag) Done() {
	c.M.Lock()
	c.Flag = 0
	c.M.Unlock()
}

type CloseFlag struct {
	Flag int
	M    sync.RWMutex
}

func (c *CloseFlag) Load() int {
	c.M.RLock()
	ret := c.Flag
	c.M.RUnlock()
	return ret
}

func (c *CloseFlag) Close() {
	c.M.Lock()
	c.Flag = 1
	c.M.Unlock()
}

type SendProcess struct {
	Num int
	M   sync.Mutex
}

func (p *SendProcess) Add() {
	p.M.Lock()
	p.Num++
	p.M.Unlock()
}

func (p *SendProcess) Done() {
	p.M.Lock()
	p.Num--
	p.M.Unlock()
}

type PullProcess struct {
	Num int
	M   sync.Mutex
}

func (p *PullProcess) Add() {
	p.M.Lock()
	p.Num++
	p.M.Unlock()
}

func (p *PullProcess) Done() {
	p.M.Lock()
	p.Num--
	p.M.Unlock()
}

type History struct {
	H []interface{}
	M sync.Mutex
}

func (h *History) Add(message interface{}) {
	h.M.Lock()
	h.H = append(h.H, message)
	h.M.Unlock()
}
