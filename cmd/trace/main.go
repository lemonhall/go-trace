package main

import "github.com/tj/go-trace/plugins/live"
import "github.com/tj/docopt"
import "encoding/json"
import "log"
import "os"

var Version = "0.0.1"

const Usage = `
  Usage:
    trace <name>
    trace -h | --help
    trace --version

  Options:
    -h, --help       output help information
    -v, --version    output version

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	if err != nil {
		log.Fatalf("failed to parse arguments: %s", err)
	}

	name := args["<name>"].(string)

	c, err := live.Dial(name)
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer c.Close()

	enc := json.NewEncoder(os.Stdout)
	for c.Next() {
		err := enc.Encode(c.Event)
		if err != nil {
			log.Fatalf("failed to write: %s", err)
		}
	}

	if c.Error != nil {
		log.Fatalf("failed to read: %s", c.Error)
	}
}
