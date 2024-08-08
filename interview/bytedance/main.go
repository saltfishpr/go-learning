package main

func main() {
	count := 10
	done := make(chan struct{})
	w1 := NewWorker(done)
	w2 := NewWorker(done)
	w3 := NewWorker(done)
	go w1.Work(func() {
		println(1)
	}, w2)
	go w2.Work(func() {
		println(2)
	}, w3)
	go w3.Work(func() {
		println(3)
		if w3.count > count {
			close(done)
		}
	}, w1)
	w1.Notify()
	<-done
}

type Worker struct {
	sig   chan struct{}
	done  chan struct{}
	count int
}

func NewWorker(done chan struct{}) *Worker {
	return &Worker{
		sig:   make(chan struct{}),
		done:  done,
		count: 0,
	}
}

func (w *Worker) Notify() {
	w.sig <- struct{}{}
}

func (w *Worker) Work(f func(), next *Worker) {
	for {
		select {
		case <-w.done:
			close(w.sig)
			return
		case <-w.sig:
			w.count++
			f()
			next.Notify()
		}
	}
}
