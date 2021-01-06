package pkg

type Bucket struct {
	Points []float64
	Count  int64
	next   *Bucket
}

func (b *Bucket) Append(val float64) {
	b.Points = append(b.Points, val)
	b.Count++
}

func (b *Bucket) Add(offset int, val float64) {
	b.Points[offset] += val
	b.Count++
}

func (b *Bucket) Reset() {
	b.Points = b.Points[:0]
	b.Count = 0
}

func (b *Bucket) Next() *Bucket {
	return b.next
}
