package trace

// import "github.com/bmizerany/assert"
import "testing"
import "bytes"

type Event struct {
	Some string
	Data string
}

func TestEmit(t *testing.T) {
	w := bytes.NewBuffer(nil)

	trace := New()
	trace.Use(writer{w})

	trace.Emit(Event{"foo", "bar"})
	trace.Emit(Event{"bar", "baz"})

	println(string(w.Bytes()))
}
