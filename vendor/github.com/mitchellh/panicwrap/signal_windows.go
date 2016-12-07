// +build windows

package panicwrap

import (
	"os"
)

var forwardSignals []os.Signal = []os.Signal{os.Interrupt}
