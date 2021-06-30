package tap

// Tap is an idiom for a water tap.
type Tap struct {
	ch   chan struct{}
	open bool
}

// New returns a new closed tap.
func New() *Tap {
	tap := &Tap{ch: make(chan struct{})}
	return tap
}

// NewOpen returns a new open tap.
func NewOpen() *Tap {
	tap := New()
	tap.Open()
	return tap
}

// Open the tap, let the stream flow.
func (t *Tap) Open() {
	if t.open {
		return
	}
	t.open = true
	close(t.ch)
}

// Close the tap, stop the stream flow.
func (t *Tap) Close() {
	if !t.open {
		return
	}
	t.open = false
	t.ch = make(chan struct{})
}

// Stream flows when the tap is open.
func (t *Tap) Stream() <-chan struct{} {
	return t.ch
}
