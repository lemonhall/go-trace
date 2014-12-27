package live

import "encoding/json"
import "time"
import "net"
import "fmt"

// Client connects to a trace socket for remote reporting.
type Client struct {
	Event map[string]interface{}
	Error error
	name  string
	dec   *json.Decoder
	net.Conn
}

// Dial connects to trace socket `name`.
func Dial(name string) (*Client, error) {
	c := &Client{
		name: name,
	}

	conn, err := net.DialTimeout("unix", c.Path(), 10*time.Second)
	if err != nil {
		return nil, err
	}

	c.Conn = conn
	c.dec = json.NewDecoder(conn)
	return c, nil
}

// Path to socket.
func (c *Client) Path() string {
	return fmt.Sprintf("/tmp/%s", c.name)
}

// Next event.
func (c *Client) Next() bool {
	c.Error = c.dec.Decode(&c.Event)
	return c.Error == nil
}
