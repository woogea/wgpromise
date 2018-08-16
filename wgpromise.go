package wgpromise

// Promise struct
type Promise struct {
	resolve func()
	reject  func()
}

func newPromise(resolve func(), reject func()) *Promise {
	p := new(Promise)
	p.reject = reject
	p.resolve = resolve
	return p
}

// Run promise run
func (p *Promise) Run(task func() bool) *Promise {
	if task() {
		p.resolve()
	} else {
		p.reject()
	}
	return p
}
