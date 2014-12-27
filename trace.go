package trace

import "encoding/json"
import "log"
import "os"
import "io"

// New tracer.
type Tracer struct {
	events  chan interface{}
	plugins []Plugin
}

// New tracer.
func New() *Tracer {
	t := &Tracer{
		events: make(chan interface{}),
	}

	go t.start()

	return t
}

// Use plugin to receive events.
func (t *Tracer) Use(plugin Plugin) {
	t.plugins = append(t.plugins, plugin)
}

// Emit event.
func (t *Tracer) Emit(e interface{}) {
	t.events <- e
}

// start relaying events.
func (t *Tracer) start() {
	for e := range t.events {
		for _, plugin := range t.plugins {
			plugin.HandleEvent(e)
		}
	}
}

// Plugin interface for handling events.
type Plugin interface {
	Name() string
	HandleEvent(e interface{})
}

// stdio writer plugin.
type stdio struct {
	io.Writer
}

// Stdio plugin outputting JSON to stderr.
var Stdio = stdio{os.Stderr}

// Name implementation.
func (p stdio) Name() string {
	return "stdio"
}

// HandleEvent implementation.
func (p stdio) HandleEvent(e interface{}) {
	err := json.NewEncoder(p).Encode(e)
	if err != nil {
		log.Printf("trace: failed to encode json: %s", err)
	}
}
