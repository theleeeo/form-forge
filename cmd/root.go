package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

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

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("Could not read config file, error:", err)
		}
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Could not find home directory, error:", err)
		}

		// check if the config file exists
		_, err = os.Stat(filepath.Join(home, ".thor.yml"))
		if err == nil {
			viper.AddConfigPath(home)
			viper.SetConfigType("yml")
			viper.SetConfigName(".thor")

			if err := viper.ReadInConfig(); err != nil {
				log.Fatal("Could not read config file, error:", err)
			}
		}
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()
}
