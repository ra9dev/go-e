package receiver

import (
	"context"
	"fmt"
	"github.com/ra9dev/go-e/pkg/timer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

// TimeConsumer consumes time from timer.TimerServer
type TimeConsumer struct {
	addr              string
	requestsFrequency time.Duration

	conn   *grpc.ClientConn
	client timer.TimerClient
}

// NewTimeConsumer constructor
func NewTimeConsumer(config *TimeConsumerConfig) *TimeConsumer {
	consumer := &TimeConsumer{
		addr:              config.APIAddr,
		requestsFrequency: config.RequestsFrequency,
	}

	return consumer
}

// Run periodically requests CurrentTime from timer.TimerServer
func (c *TimeConsumer) Run(ctx context.Context) error {
	if err := c.setupClient(); err != nil {
		return err
	}

	ticker := time.NewTicker(c.requestsFrequency)
	for {
		select {
		case <-ctx.Done():
			return c.conn.Close()
		case <-ticker.C:
			timeStamp, err := c.client.CurrentTime(ctx, &emptypb.Empty{})
			if err != nil {
				log.Printf("[ERROR] could not request time: %v", err)
				continue
			}

			log.Printf("Received %s", timeStamp.AsTime().String())
		}
	}
}

func (c *TimeConsumer) setupClient() error {
	log.Printf("[INFO] Connecting to %s", c.addr)
	conn, err := grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("could not connect to api: %w", err)
	}

	c.conn = conn
	c.client = timer.NewTimerClient(conn)
	return nil
}
