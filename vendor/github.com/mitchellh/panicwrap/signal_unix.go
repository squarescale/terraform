// +build !windows

package panicwrap

import (
	"os"
	"syscall"
)

var forwardSignals []os.Signal = []os.Signal{os.Interrupt, syscall.SIGTERM}
