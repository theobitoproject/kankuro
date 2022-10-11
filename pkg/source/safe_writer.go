package source

import (
	"io"
	"sync"
)

// safeWriter allows to write bytes to the underlying data stream in a safe manner
type safeWriter struct {
	w  io.Writer
	mu sync.Mutex
}

// newSafeWriter creates a new instance of SafeWriter
func newSafeWriter(w io.Writer) io.Writer {
	return &safeWriter{
		w: w,
	}
}

// Write writes len(p) bytes from p to the underlying data stream
func (sw *safeWriter) Write(p []byte) (int, error) {
	sw.mu.Lock()
	defer sw.mu.Unlock()
	return sw.w.Write(p)
}
