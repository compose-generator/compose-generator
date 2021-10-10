/*
Copyright 2021 Compose Generator Contributors
All rights reserved ©
*/
package logrus

import (
	"bytes"
	"sync"
)

var (
	bufferPool BufferPool
)

type BufferPool interface {
	Put(*bytes.Buffer)
	Get() *bytes.Buffer
}

type defaultPool struct {
	pool *sync.Pool
}

func (p *defaultPool) Put(buf *bytes.Buffer) {
	p.pool.Put(buf)
}

func (p *defaultPool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

func getBuffer() *bytes.Buffer {
	return bufferPool.Get()
}

func putBuffer(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

// SetBufferPool allows to replace the default logrus buffer pool
// to better meets the specific needs of an application.
func SetBufferPool(bp BufferPool) {
	bufferPool = bp
}

func init() {
	SetBufferPool(&defaultPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	})
}
