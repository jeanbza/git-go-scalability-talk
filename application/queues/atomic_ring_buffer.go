package queues

import (
	"sync/atomic"
	"unsafe"
)

type bucket struct {
	data []byte
	seq  uint64
}

type OneToOne struct {
	buffer     []unsafe.Pointer
	writeIndex uint64
	readIndex  uint64
}

// This is only suitable for one-to-one usage
func NewAtomicRingBuffer(size int) *OneToOne {
	d := &OneToOne{
		buffer: make([]unsafe.Pointer, size),
	}

	d.writeIndex = ^d.writeIndex
	return d
}

func (d *OneToOne) Enqueue(data []byte) {
	writeIndex := atomic.AddUint64(&d.writeIndex, 1)
	idx := writeIndex % uint64(len(d.buffer))
	newBucket := &bucket{
		data: data,
		seq:  writeIndex,
	}

	atomic.StorePointer(&d.buffer[idx], unsafe.Pointer(newBucket))
}

func (d *OneToOne) Dequeue() ([]byte, bool) {
	readIndex := atomic.LoadUint64(&d.readIndex)
	idx := readIndex % uint64(len(d.buffer))

	value, ok := d.tryNext(idx)
	if ok {
		atomic.AddUint64(&d.readIndex, 1)
	}
	return value, ok
}

func (d *OneToOne) tryNext(idx uint64) ([]byte, bool) {
	result := (*bucket)(atomic.SwapPointer(&d.buffer[idx], nil))

	if result == nil {
		return nil, false
	}

	if result.seq > d.readIndex {
		atomic.StoreUint64(&d.readIndex, result.seq)
	}

	return result.data, true
}
