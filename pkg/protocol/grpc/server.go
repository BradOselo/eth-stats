package grpc

import (
	"context"

	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"

	"github.com/cardenasrjl/eth-stats/pkg/api/v1"
	"github.com/cardenasrjl/eth-stats/pkg/logger"
	"github.com/cardenasrjl/eth-stats/pkg/protocol/grpc/middleware"
)

// RunServer runs gRPC service to publish 
func RunServer(ctx context.Context, v1APIFeeService v1.FeeServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}


	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)


	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterFeeServiceServer(server, v1APIFeeService)
	

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}