package natssrl

import (
	"fmt"

	"github.com/go-logr/logr"
	"github.com/nats-io/nats.go"
	"github.com/netw-device-driver/netwdevpb"
	"google.golang.org/protobuf/proto"
)

// Client holds the state of the Nats client
type Client struct {
	Server string
	Topic  string
	Log    logr.Logger
}

// Publish publishes the topic/message to the NATS server
func (c *Client) Publish(d *netwdevpb.ConfigMessage) error {
	// Connect Options.
	opts := []nats.Option{nats.Name(fmt.Sprintf("NATS Publisher %s", c.Topic))}

	// Connect to NATS
	nc, err := nats.Connect(c.Server, opts...)
	if err != nil {
		return fmt.Errorf("Nats connect error: %s", err)
	}
	defer nc.Close()

	data, err := proto.Marshal(d)
	if err != nil {
		return fmt.Errorf("proto Marshal error: %s", err)
	}

	nc.Publish(c.Topic, data)
	nc.Flush()
	if err := nc.LastError(); err != nil {
		return fmt.Errorf("Nats publish error: %s", err)
	}

	return nil
}
