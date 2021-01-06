package pkg

import (
	"fmt"
	"sync"
	"time"
)

type RollingCounter struct {
	mu     sync.RWMutex
	window *Window
	size   int
	offset int

	bucketDuration time.Duration
	lastAppendTime time.Time
}

type RollingCounterOpts struct {
	Size           int
	BucketDuration time.Duration
}

func NewRollingCounter(opts RollingCounterOpts) *RollingCounter {
	window := NewWindow(WindowOpts{Size: opts.Size})
	return &RollingCounter{
		//mu:             sync.RWMutex{},
		window:         window,
		size:           window.Size(),
		offset:         0,
		bucketDuration: opts.BucketDuration,
		lastAppendTime: time.Now(),
	}
}

func (r *RollingCounter) Timespan() int {
	v := int(time.Since(r.lastAppendTime) / r.bucketDuration)
	if v > -1 { // maybe time backwards
		return v
	}
	return r.size
}

func (r *RollingCounter) add(f func(offset int, val float64), val float64) {
	r.mu.Lock()
	timespan := r.Timespan()
	if timespan > 0 {
		r.lastAppendTime = r.lastAppendTime.Add(time.Duration(timespan * int(r.bucketDuration)))
		offset := r.offset
		// reset the expired buckets
		s := offset + 1
		if timespan > r.size {
			timespan = r.size
		}
		e, e1 := s+timespan, 0 // e: reset offset must start from offset+1
		if e > r.size {
			e1 = e - r.size
			e = r.size
		}
		for i := s; i < e; i++ {
			r.window.ResetBucket(i)
			offset = i
		}
		for i := 0; i < e1; i++ {
			r.window.ResetBucket(i)
			offset = i
		}
		r.offset = offset
	}
	f(r.offset, val)
	r.mu.Unlock()
}

func (r *RollingCounter) Append(val float64) {
	r.add(r.window.Append, val)
}

func (r *RollingCounter) Add(val float64) {
	if val < 0 {
		panic(fmt.Errorf("stat/metric: cannot decrease in value. val: %d", val))
	}
	r.add(r.window.Add, val)
}

func (r *RollingCounter) Reduce(f func(Iterator) float64) (val float64) {
	r.mu.RLock()
	timespan := r.Timespan()
	if count := r.size - timespan; count > 0 {
		offset := r.offset + timespan + 1
		if offset >= r.size {
			offset = offset - r.size
		}
		val = f(r.window.Iterator(offset, count))
	}
	r.mu.RUnlock()
	return val
}


func (r *RollingCounter) Avg() float64 {
	return r.Reduce(Avg)
}

func (r *RollingCounter) Min() float64 {
	return r.Reduce(Min)
}

func (r *RollingCounter) Max() float64 {
	return r.Reduce(Max)
}

func (r *RollingCounter) Sum() float64 {
	return r.Reduce(Sum)
}

func (r *RollingCounter) Value() int64 {
	return int64(r.Sum())
}
