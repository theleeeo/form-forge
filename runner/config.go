package runner

type Config struct {
	GrpcAddress string   `yaml:"grpc-address"`
	HttpAddress string   `yaml:"http-address"`
	RepoCfg     PgConfig `yaml:"repo"`
}

type PgConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
