package runner

type Config struct {
	GrpcAddress string      `yaml:"grpc-address"`
	HttpAddress string      `yaml:"http-address"`
	RepoCfg     MySqlConfig `yaml:"repo"`
}

type MySqlConfig struct {
	Address  string `yaml:"address"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
