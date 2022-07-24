package empty_struct

var empty = struct{}{}

type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

func (s Semaphore) P(n int) {
	for i := 0; i < n; i++ {
		s <- empty
	}
}

func (s Semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
