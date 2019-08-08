package jar

type Bucket struct {
	burst   int
	content chan int
}

func (b *Bucket) Give() {
	b.content <- 1
}

func (b *Bucket) Take() {
	<-b.content
}

func (b *Bucket) Full() bool {
	return len(b.content) == b.burst
}

func NewBucket(burst int) Jar {
	content := make(chan int, burst)
	return &Bucket{
		burst:   burst,
		content: content,
	}
}
