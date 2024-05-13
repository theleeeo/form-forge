package runner

import "github.com/theleeeo/form-forge/repo"

type Config struct {
	GrpcAddress string           `yaml:"grpc-address"`
	HttpAddress string           `yaml:"http-address"`
	RepoCfg     repo.MySqlConfig `yaml:"repo"`
}
