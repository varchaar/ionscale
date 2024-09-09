package term

import (
	"github.com/jsiebens/go-edit/interrupt"
	"github.com/moby/term"
	"io"
	"os"
)

// SafeFunc is a function to be invoked by TTY.
type SafeFunc func() error

type TTY struct {
	// In is a reader representing stdin. It is a required field.
	In io.Reader

	// TryDev indicates the TTY should try to open /dev/tty if the provided input
	// is not a file descriptor.
	TryDev bool
}

func (t TTY) Safe(fn SafeFunc) error {
	inFd, isTerminal := term.GetFdInfo(t.In)

	if !isTerminal && t.TryDev {
		if f, err := os.Open("/dev/tty"); err == nil {
			defer f.Close()
			inFd = f.Fd()
			isTerminal = term.IsTerminal(inFd)
		}
	}
	if !isTerminal {
		return fn()
	}

	state, err := term.SaveState(inFd)
	if err != nil {
		return err
	}

	return interrupt.Chain(nil, func() {
		term.RestoreTerminal(inFd, state)
	}).Run(fn)
}
