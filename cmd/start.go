package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/theleeeo/form-forge/runner"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// Todo: Do not log secrets
		log.Println("Config:", prettyPrint(cfg))

		if err := runner.Run(cfg); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	},
}

func prettyPrint(i interface{}) string {
	s, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}
