package client

import (
	"codeberg.org/kasefuchs/go-kit/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Target string             `koanf:"target"`
	TLS    *cert.ClientConfig `koanf:"tls"`
}

func (c Config) DialOptions() ([]grpc.DialOption, error) {
	opts := make([]grpc.DialOption, 0, 1)

	if c.TLS == nil {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		tlsCfg, err := c.TLS.TLSConfig()
		if err != nil {
			return nil, err
		}

		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)))
	}

	return opts, nil
}
