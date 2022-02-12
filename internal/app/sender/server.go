package sender

import (
	"context"
	"fmt"
	timer "github.com/ra9dev/go-e/pkg/timer"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"sync"
	"time"
)

// TimeSender sends current time
type TimeSender struct {
	addr string

	grpcServer *grpc.Server
	timer.UnimplementedTimerServer
}

func NewTimeSender(config *TimeSenderConfig) *TimeSender {
	return &TimeSender{
		addr: config.ListenAddr,
	}
}

func (s *TimeSender) CurrentTime(context.Context, *emptypb.Empty) (*timestamppb.Timestamp, error) {
	now := time.Now().UTC()
	log.Printf("Sending %s", now.String())

	return timestamppb.New(now), nil
}

func (s *TimeSender) Run(ctx context.Context) error {
	listener, err := s.setupServer()
	if err != nil {
		return fmt.Errorf("could not setup server: %v", err)
	}

	wg := new(sync.WaitGroup)
	// 2 for server and graceful shutdown
	wg.Add(2)

	go s.serve(listener, wg)
	go s.gracefulShutdown(ctx, wg)

	wg.Wait()
	return nil
}

func (s *TimeSender) setupServer() (net.Listener, error) {
	s.grpcServer = grpc.NewServer()
	timer.RegisterTimerServer(s.grpcServer, s)

	return net.Listen("tcp", s.addr)
}

func (s *TimeSender) serve(listener net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("[INFO] serving GRPC on \"%s\"", s.addr)
	if err := s.grpcServer.Serve(listener); err != nil {
		log.Fatalf("[ERROR] grpc serve: %v", err)
	}
}

func (s *TimeSender) gracefulShutdown(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	<-ctx.Done()
	log.Printf("[INFO] shutting down gRPC server")
	s.grpcServer.GracefulStop()
}
