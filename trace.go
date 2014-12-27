package trace

// Plugin interface.
type Plugin interface {
	HandleEvent(e interface{})
}

// Tracer.
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
