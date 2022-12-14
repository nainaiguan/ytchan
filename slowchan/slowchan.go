package slowchan

import (
	"context"
	"sync"
	"time"
)

type slowChan struct {
	data           chan interface{}
	cap            int
	step           time.Duration
	maxSendProcess int
	sendHistory    history
	sendProcess    sendProcess
	cleanFlag      cleanFlag
	closeFlag      closeFlag
	ctx            context.Context
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
	h []interface{}
	m sync.RWMutex
}

func (h *history) Load() []interface{} {
	h.m.RLock()
	ret := h.h
	h.m.RUnlock()
	return ret
}

func (h *history) Add(message interface{}) {
	h.m.Lock()
	h.h = append(h.h, message)
	h.m.Unlock()
}

func (h *history) Len() int {
	h.m.RLock()
	defer h.m.RUnlock()
	return len(h.h)
}
