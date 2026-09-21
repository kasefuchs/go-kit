package cert

import "crypto/tls"

type Config struct {
	CertFile string `koanf:"cert_file"`
	KeyFile  string `koanf:"key_file"`
}

func (c Config) TLSConfig() (*tls.Config, error) {
	cfg := &tls.Config{}

	if c.CertFile != "" || c.KeyFile != "" {
		cert, err := LoadKeyPair(c.CertFile, c.KeyFile)
		if err != nil {
			return nil, err
		}

		cfg.Certificates = []tls.Certificate{cert}
	}

	return cfg, nil
}

type ClientConfig struct {
	Config `koanf:",squash"`

	CAFile             string `koanf:"ca_file"`
	ServerName         string `koanf:"server_name"`
	InsecureSkipVerify bool   `koanf:"insecure_skip_verify"`
}

func (c ClientConfig) TLSConfig() (*tls.Config, error) {
	cfg, err := c.Config.TLSConfig()
	if err != nil {
		return nil, err
	}

	cfg.ServerName = c.ServerName
	cfg.InsecureSkipVerify = c.InsecureSkipVerify

	if c.CAFile != "" {
		pool, err := LoadCA(c.CAFile)
		if err != nil {
			return nil, err
		}

		cfg.RootCAs = pool
	}

	return cfg, nil
}
