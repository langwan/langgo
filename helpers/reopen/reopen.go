package reopen

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Reopener interface {
	Reopen() error
}

type Writer interface {
	Reopener
	io.Writer
}

type WriteCloser interface {
	Reopener
	io.WriteCloser
}

type FileWriter struct {
	mu   sync.Mutex
	f    *os.File
	mode os.FileMode
	name string
}

func (f *FileWriter) Fd() uintptr {
	f.mu.Lock()
	defer f.mu.Unlock()
	ret := f.f.Fd()
	return ret
}

func (f *FileWriter) Close() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	err := f.f.Close()
	return err
}

// mutex free version
func (f *FileWriter) reopen() error {
	if f.f != nil {
		f.f.Close()
		f.f = nil
	}
	newf, err := os.OpenFile(f.name, os.O_WRONLY|os.O_APPEND|os.O_CREATE, f.mode)
	if err != nil {
		f.f = nil
		return err
	}
	f.f = newf

	return nil
}

// Reopen the file
func (f *FileWriter) Reopen() error {
	f.mu.Lock()
	defer f.mu.Unlock()
	err := f.reopen()
	return err
}

// Write implements the stander io.Writer interface
func (f *FileWriter) Write(p []byte) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	n, err := f.f.Write(p)
	return n, err
}

func NewFileWriter(name string) (*FileWriter, error) {
	return NewFileWriterMode(name, 0666)
}

func NewFileWriterMode(name string, mode os.FileMode) (*FileWriter, error) {
	writer := FileWriter{
		f:    nil,
		name: name,
		mode: mode,
	}
	err := writer.reopen()
	if err != nil {
		return nil, err
	}
	return &writer, nil
}

type BufferedFileWriter struct {
	mu         sync.Mutex
	quitChan   chan bool
	done       bool
	origWriter *FileWriter
	bufWriter  *bufio.Writer
}

func (bw *BufferedFileWriter) Reopen() error {
	bw.mu.Lock()
	bw.bufWriter.Flush()

	// use non-mutex version since we are using this one
	err := bw.origWriter.reopen()

	bw.bufWriter.Reset(io.Writer(bw.origWriter))
	bw.mu.Unlock()

	return err
}

func (bw *BufferedFileWriter) Close() error {
	bw.quitChan <- true
	bw.mu.Lock()
	bw.done = true
	bw.bufWriter.Flush()
	bw.origWriter.f.Close()
	bw.mu.Unlock()
	return nil
}

func (bw *BufferedFileWriter) Write(p []byte) (int, error) {
	bw.mu.Lock()
	n, err := bw.bufWriter.Write(p)

	if bw.bufWriter.Buffered() < len(p) {
		bw.bufWriter.Flush()
	}

	bw.mu.Unlock()
	return n, err
}

func (bw *BufferedFileWriter) Flush() {
	bw.mu.Lock()

	bw.bufWriter.Flush()
	bw.origWriter.f.Sync()
	bw.mu.Unlock()
}

func (bw *BufferedFileWriter) flushDaemon(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-bw.quitChan:
			ticker.Stop()
			return
		case <-ticker.C:
			bw.Flush()
		}
	}
}

const bufferSize = 256 * 1024
const flushInterval = 30 * time.Second

func NewBufferedFileWriter(w *FileWriter) *BufferedFileWriter {
	return NewBufferedFileWriterSize(w, bufferSize, flushInterval)
}

func NewBufferedFileWriterSize(w *FileWriter, size int, flush time.Duration) *BufferedFileWriter {
	bw := BufferedFileWriter{
		quitChan:   make(chan bool, 1),
		origWriter: w,
		bufWriter:  bufio.NewWriterSize(w, size),
	}
	go bw.flushDaemon(flush)
	return &bw
}

type multiReopenWriter struct {
	writers []Writer
}

func (t *multiReopenWriter) Reopen() error {
	for _, w := range t.writers {
		err := w.Reopen()
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *multiReopenWriter) Write(p []byte) (int, error) {
	for _, w := range t.writers {
		n, err := w.Write(p)
		if err != nil {
			return n, err
		}
		if n != len(p) {
			return n, io.ErrShortWrite
		}
	}
	return len(p), nil
}

func MultiWriter(writers ...Writer) Writer {
	w := make([]Writer, len(writers))
	copy(w, writers)
	return &multiReopenWriter{w}
}

type nopReopenWriteCloser struct {
	io.Writer
}

func (nopReopenWriteCloser) Reopen() error {
	return nil
}

func (nopReopenWriteCloser) Close() error {
	return nil
}

func NopWriter(w io.Writer) WriteCloser {
	return nopReopenWriteCloser{w}
}

var (
	Stdout  = NopWriter(os.Stdout)
	Stderr  = NopWriter(os.Stderr)
	Discard = NopWriter(ioutil.Discard)
)
