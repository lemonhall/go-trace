package main

import "github.com/tj/go-trace/plugins/live"
import "github.com/tj/docopt"
import "text/template"
import "encoding/json"
import "log"
import "os"

var Version = "0.0.1"

const Usage = `
  Usage:
    trace <name> [--format tmpl]
    trace -h | --help
    trace --version

  Options:
    -f, --format tmpl   output template
    -h, --help          output help information
    -v, --version       output version

`

func main() {
	args, err := docopt.Parse(Usage, nil, true, Version, false)
	if err != nil {
		log.Fatalf("failed to parse arguments: %s", err)
	}

	name := args["<name>"].(string)
	out := identity

	if s, ok := args["--format"].(string); ok {
		tmpl := template.Must(template.New("format").Parse(s))
		out = func(e interface{}) error {
			err := tmpl.Execute(os.Stdout, e)
			if err != nil {
				return err
			}

			_, err = os.Stdout.WriteString("\n")
			return err
		}
	}

	c, err := live.Dial(name)
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer c.Close()

	for c.Next() {
		err := out(c.Event)
		if err != nil {
			log.Fatalf("failed to write: %s", err)
		}
	}

	if c.Error != nil {
		log.Fatalf("failed to read: %s", c.Error)
	}
}

func identity(e interface{}) error {
	return json.NewEncoder(os.Stdout).Encode(e)
}
