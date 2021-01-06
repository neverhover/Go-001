package pkg

type Window struct {
	buckets []Bucket
	size    int
}

type WindowOpts struct {
	Size int
}

func NewWindow(opts WindowOpts) *Window {
	buckets := make([]Bucket, opts.Size)
	for offset := range buckets {
		buckets[offset] = Bucket{Points: make([]float64, 0)}
		nextOffset := offset + 1
		if nextOffset == opts.Size {
			nextOffset = 0
		}
		buckets[offset].next = &buckets[nextOffset]
	}
	return &Window{
		buckets: buckets,
		size:    opts.Size,
	}
}

func (w *Window) Reset() {
	for offset := range w.buckets {
		w.ResetBucket(offset)
	}
}

func (w *Window) ResetBucket(offset int) {
	w.buckets[offset].Reset()
}

func (w *Window) ResetBuckets(offsets []int) {
	for _, offset := range offsets {
		w.ResetBucket(offset)
	}
}

func (w *Window) Append(offset int, val float64) {
	w.buckets[offset].Append(val)
}

func (w *Window) Add(offset int, val float64) {
	if w.buckets[offset].Count == 0 {
		w.buckets[offset].Append(val)
		return
	}
	w.buckets[offset].Add(0, val)
}

func (w *Window) Bucket(offset int) Bucket {
	return w.buckets[offset]
}

func (w *Window) Size() int {
	return w.size
}

func (w *Window) Iterator(offset int, count int) Iterator {
	return Iterator{
		count: count,
		cur:   &w.buckets[offset],
	}
}
