package trace

import "github.com/bmizerany/assert"
import "./plugins/json"
import "testing"
import "bytes"

type Event struct {
	Some string
	Data string
}

func TestEmit(t *testing.T) {
	buf := bytes.NewBuffer(nil)

	trace := New()
	trace.Use(json.New(buf))

	trace.Emit(Event{"one", "1"})
	trace.Emit(Event{"two", "2"})
	trace.Emit(Event{"three", "3"})

	exp := `{"Some":"one","Data":"1"}
{"Some":"two","Data":"2"}
{"Some":"three","Data":"3"}
`

	assert.Equal(t, exp, string(buf.Bytes()))
}

func BenchmarkEmitNoPlugins(b *testing.B) {
	t := New()
	for i := 0; i < b.N; i++ {
		t.Emit(0)
	}
}
