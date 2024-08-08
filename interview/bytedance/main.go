package main

func main() {
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
		w3.count++
		if w3.count > 10 {
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

func (p *Worker) Notify() {
	p.sig <- struct{}{}
}

func (p *Worker) Work(f func(), next *Worker) {
	for {
		select {
		case <-p.done:
			close(p.sig)
			return
		case <-p.sig:
			f()
			next.Notify()
		}
	}
}
