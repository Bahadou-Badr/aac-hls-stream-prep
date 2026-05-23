package worker

type Pool struct {
	Jobs chan Job
}

func NewPool(bufferSize int) *Pool {
	return &Pool{
		Jobs: make(chan Job, bufferSize),
	}
}
