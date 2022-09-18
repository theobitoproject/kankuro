package writers

import (
	"io"
	"sync"
)

// SafeWriter allows to write bytes to the underlying data stream in a safe manner
type SafeWriter struct {
	w  io.Writer
	mu sync.Mutex
}

// NewSafeWriter creates a new instance of SafeWriter
func NewSafeWriter(w io.Writer) io.Writer {
	return &SafeWriter{
		w: w,
	}
}

// Write writes len(p) bytes from p to the underlying data stream
func (sw *SafeWriter) Write(p []byte) (int, error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	return sw.w.Write(p)
}
