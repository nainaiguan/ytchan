package subchan

import (
	"context"
	"sync"
	"ytChan/util/prettylog"
)

type SubChan struct {
	data           chan interface{}
	cap            int
	maxSendProcess int
	subscriber     subscriber
	sendHistory    history
	sendProcess    sendProcess
	cleanFlag      cleanFlag
	closeFlag      closeFlag
	ctx            context.Context
}

type subscriber struct {
	m  map[string]chan interface{}
	mu sync.RWMutex
}

func (s *subscriber) Load(name string) chan interface{} {
	s.mu.RLock()
	if _, ok := s.m[name]; !ok {
		prettylog.Errorf("subscriber.Load Error: %s", "no such subscriber")
		s.mu.RUnlock()
		return nil
	}
	t := s.m[name]
	s.mu.RUnlock()
	return t
}

func (s *subscriber) Add(name string, size int) {
	if s.Load(name) == nil {
		prettylog.Errorf("subscriber.Add Error: %s", "duplicate subscriber")
		return
	}
	s.mu.Lock()
	m := make(chan interface{}, size)
	s.m[name] = m
	s.mu.Unlock()
}

func (s *subscriber) Drop(name string) {
	if s.Load(name) == nil {
		prettylog.Errorf("subscriber.Drop Error: %s", "no such subscriber")
		return
	}
	s.mu.Lock()
	delete(s.m, name)
	s.mu.Unlock()
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
