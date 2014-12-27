package main

import "github.com/tj/go-trace/plugins/live"
import "github.com/tj/go-trace/plugins/json"
import "github.com/tj/go-trace"
import "math/rand"
import "time"

type Trace struct {
	Name   string    `json:"name"`
	Start  time.Time `json:"start"`
	Finish time.Time `json:"finish"`
	Path   string    `json:"path"`
}

func main() {
	t := trace.New()
	t.Use(json.Stdio)
	t.Use(live.New("example"))

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			start := time.Now()
			t.Emit(&Trace{"start", start, start, "/foo"})
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			t.Emit(&Trace{"finish", start, time.Now(), "/foo"})
		}
	}()

	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			start := time.Now()
			t.Emit(&Trace{"start", start, start, "/bar"})
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			t.Emit(&Trace{"finish", start, time.Now(), "/bar"})
		}
	}()

	select {}
}
