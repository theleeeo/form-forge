package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theleeeo/form-forge/runner"
)

var cfgFile string

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ~/.cfg.yml)")

	rootCmd.AddCommand(startCmd)
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "formforge",
	Short: "A Form Service",
}

func loadConfig() (runner.Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("cfg")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.formforge")
		viper.AddConfigPath("/etc/formforge/")
	}

	if err := viper.ReadInConfig(); err != nil {
		var notFoundErr viper.ConfigFileNotFoundError
		if errors.As(err, &notFoundErr) {
			return runner.Config{}, errors.New("config file not found")
		}

		return runner.Config{}, fmt.Errorf("error reading config file: %w", err)
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	cfg := runner.Config{
		Addr: viper.GetString("addr"),
		RepoCfg: runner.PgConfig{
			Host:     viper.GetString("repo.host"),
			Port:     viper.GetInt("repo.port"),
			User:     viper.GetString("repo.user"),
			Password: viper.GetString("repo.password"),
			Database: viper.GetString("repo.database"),
		},
	}

	if err := cfg.Validate(); err != nil {
		return runner.Config{}, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}
