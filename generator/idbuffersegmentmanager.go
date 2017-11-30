package generator


import(
	"errors"
	"log"
	"io"
	"sync"
)


type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func () (io.Closer, error)
	closed    bool  
}

var ErrPoolClosed = errors.New("Pool has been closed.")

//New create a pool for managering resources
//Size the num of resources
func New(fn func() (io.Closer, error), size uint64)(*Pool, error){
	if size <= 0 {
		return nil, errors.New("Size value must be > 0")
	}
	return &Pool{
		factory: fn,
		resources: make(chan io.Closer, size)
	}
}

func (p *Pool) Acquire() (io.Closer, error){
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
	default:
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
	default:
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	//pools closed	
	p.closed = true

	//shutdown this channel
	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}



