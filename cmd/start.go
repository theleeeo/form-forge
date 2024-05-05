package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
	"go.leeeo.se/form-forge/runner"
	"gopkg.in/yaml.v3"
)

func loadConfig() *runner.Config {
	content, err := os.ReadFile("./.cfg.yml")
	if err != nil {
		log.Fatal(err)
	}

	var config runner.Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &config
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg := loadConfig()

		log.Println("Config:", prettyPrint(cfg))

		if err := runner.Run(cfg); err != nil {
			log.Println(err)
			return nil
		}

		return nil
	},
}

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}
