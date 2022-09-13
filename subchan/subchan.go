package subchan

import (
	"context"
	"sync"
	"time"
)

type SubChan struct {
	data           chan interface{}
	cap            int
	maxSendProcess int
	subscriber     map[string]chan interface{}
	sendHistory    history
	sendProcess    sendProcess
	cleanFlag      cleanFlag
	closeFlag      closeFlag
	ctx            context.Context
}

func (d *SubChan) Reconcile() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			time.Sleep(100 * time.Millisecond)

			if len(d.data) != 0 {
				message := <-d.data
				for _, c := range d.subscriber {
					c <- message
				}
			}
		}
	}
}

func (d *SubChan) DftChanCleanDaemon() {
	for {
		select {
		case <-d.ctx.Done():
			return
		default:
			d.cleanFlag.Clean()
			l := len(d.sendHistory.h)
			tmp := make([]interface{}, l)
			copy(tmp, d.sendHistory.h)
			d.sendHistory.h = tmp
			d.cleanFlag.Done()

			time.Sleep(30 * time.Second)
		}
	}
}

type cleanFlag struct {
	m sync.RWMutex
}

func (c *cleanFlag) Load() {
	c.m.RLock()
}

func (c *cleanFlag) Free() {
	c.m.RUnlock()
}

func (c *cleanFlag) Clean() {
	c.m.Lock()
}

func (c *cleanFlag) Done() {
	c.m.Unlock()
}

type closeFlag struct {
	flag int
	m    sync.RWMutex
}

func (c *closeFlag) Load() int {
	c.m.RLock()
	ret := c.flag
	c.m.RUnlock()
	return ret
}

func (c *closeFlag) Close() {
	c.m.Lock()
	c.flag = 1
	c.m.Unlock()
}

type sendProcess struct {
	num int
	m   sync.RWMutex
}

func (p *sendProcess) Load() int {
	p.m.RLock()
	ret := p.num
	p.m.RUnlock()
	return ret
}

func (p *sendProcess) Add() {
	p.m.Lock()
	p.num++
	p.m.Unlock()
}

func (p *sendProcess) Done() {
	p.m.Lock()
	p.num--
	p.m.Unlock()
}

type history struct {
	max int
	h   []interface{}
	m   sync.RWMutex
}

func (h *history) Load() []interface{} {
	h.m.RLock()
	ret := h.h
	h.m.RUnlock()
	return ret
}

func (h *history) Add(message interface{}) {
	h.m.Lock()
	if len(h.h) >= h.max {
		t := append(h.h[1:], message)
		h.h = t
		return
	}
	h.h = append(h.h, message)
	h.m.Unlock()
}
