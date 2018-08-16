package wgpromise

// Promise struct
type status int

const (
	//Pending status is pending
	Pending = iota
	//Rejected status is rejected
	Rejected
	//Resolved status is resolved
	Resolved
)

// Promise is promise
type Promise struct {
	resolve func()
	reject  func()
	stat    status
}

func (p *Promise) getStatus() status {
	return p.stat
}
func newPromise(resolve func(), reject func()) *Promise {
	p := new(Promise)
	p.reject = reject
	p.resolve = resolve
	return p
}

// Run promise run
func (p *Promise) Run(task func() bool) *Promise {
	p.stat = Pending
	go func() {
		if task() {
			p.stat = Resolved
			p.resolve()
		} else {
			p.stat = Rejected
			p.reject()
		}
	}()
	return p
}
