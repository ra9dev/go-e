package sender

import (
	"context"
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

	timer.UnimplementedTimerServer
}

func (s *TimeSender) CurrentTime(context.Context, *emptypb.Empty) (*timestamppb.Timestamp, error) {
	now := time.Now().UTC()
	log.Printf("Sending %s", now.String())

	return timestamppb.New(now), nil
}

func NewTimeSender(config *TimeSenderConfig) *TimeSender {
	return &TimeSender{
		addr: config.ListenAddr,
	}
}

func (s *TimeSender) Run(ctx context.Context) {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	timer.RegisterTimerServer(grpcServer, s)

	wg := new(sync.WaitGroup)
	// 2 for server and graceful shutdown
	wg.Add(2)

	go s.serve(grpcServer, listener, wg)
	go s.gracefulShutdown(ctx, grpcServer, wg)

	wg.Wait()
}

func (s *TimeSender) serve(grpcServer *grpc.Server, listener net.Listener, wg *sync.WaitGroup) {
	defer wg.Done()

	log.Printf("[INFO] serving GRPC on \"%s\"", s.addr)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("[ERROR] grpc serve: %v", err)
	}
}

func (s *TimeSender) gracefulShutdown(ctx context.Context, grpcServer *grpc.Server, wg *sync.WaitGroup) {
	defer wg.Done()

	<-ctx.Done()
	log.Printf("[INFO] shutting down gRPC server")
	grpcServer.GracefulStop()
}
