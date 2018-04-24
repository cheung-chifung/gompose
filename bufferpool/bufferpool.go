package bufferpool

import (
	"bytes"
	"sync"
)

var bufferPool = &sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func Get() *bytes.Buffer {
	b, ok := bufferPool.Get().(*bytes.Buffer)
	if !ok {
		b = bufferPool.New().(*bytes.Buffer)
	}
	return b
}

func Free(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}
