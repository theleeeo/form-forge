package runner

import "errors"

type Config struct {
	ApiAddr    string
	PublicAddr string
	RepoCfg    PgConfig
}

func (c Config) Validate() error {
	if c.ApiAddr == "" {
		return errors.New("missing api address")
	}

	if c.PublicAddr == "" {
		return errors.New("missing public address")
	}

	if err := c.RepoCfg.Validate(); err != nil {
		return err
	}

	return nil
}

type PgConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func (c PgConfig) Validate() error {
	if c.Host == "" {
		return errors.New("missing host")
	}

	if c.Port == 0 {
		return errors.New("missing port")
	}

	if c.User == "" {
		return errors.New("missing user")
	}

	if c.Password == "" {
		return errors.New("missing password")
	}

	if c.Database == "" {
		return errors.New("missing database")
	}

	return nil
}
