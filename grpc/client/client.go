package client

import (
	"codeberg.org/kasefuchs/go-kit/log"
	"google.golang.org/grpc"
)

func New(cfg Config, extra ...grpc.DialOption) (*grpc.ClientConn, error) {
	opts, err := cfg.DialOptions()
	if err != nil {
		return nil, err
	}
	
	return grpc.NewClient(cfg.Target, append(opts, extra...)...)
}

func MustNew(cfg Config, extra ...grpc.DialOption) *grpc.ClientConn {
	conn, err := New(cfg, extra...)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create gRPC client")
	}

	return conn
}
