package json

import j "encoding/json"
import "log"
import "os"
import "io"

// writer plugin.
type writer struct {
	io.Writer
}

// Stdio plugin.
var Stdio = New(os.Stderr)

// New json writer plugin.
func New(w io.Writer) *writer {
	return &writer{w}
}

// HandleEvent implementation.
func (p writer) HandleEvent(e interface{}) {
	err := j.NewEncoder(p).Encode(e)
	if err != nil {
		log.Printf("trace: failed to encode json: %s", err)
	}
}
