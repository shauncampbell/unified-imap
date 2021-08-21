package main

import (
	"fmt"
	"os"

	"github.com/shauncampbell/unified-imap/internal/credentials"

	"github.com/shauncampbell/unified-imap/internal/authenticator"
	"github.com/shauncampbell/unified-imap/internal/config"
	"github.com/shauncampbell/unified-imap/internal/unified"

	"github.com/emersion/go-imap/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "imap",
	RunE: runApplication,
}

const configurationFileFlag string = "configuration-file"

func runApplication(cmd *cobra.Command, args []string) error {
	// Figure out the config file path
	cfgFile, err := cmd.PersistentFlags().GetString(configurationFileFlag)
	if err != nil {
		return fmt.Errorf("must provide configuration file: %w", err)
	}
	// Read in the config file
	cfg, err := config.ReadConfigFromFile(cfgFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// Create a new memory authenticator
	auth, err := authenticator.ForConfiguration(&cfg.Authenticator)
	if err != nil {
		return fmt.Errorf("failed to load authenticator: %s", err.Error())
	}

	// Create a new memory credentials store
	creds, err := credentials.ForConfiguration(&cfg.Credentials)
	if err != nil {
		return fmt.Errorf("failed to load credentials store: %s", err.Error())
	}

	// Create a memory backend
	be := unified.New(auth, creds)

	// Create a new server
	s := server.New(be)
	s.Addr = ":1143"
	// Since we will use this server for testing only, we can allow plain text
	// authentication over unencrypted connections
	s.AllowInsecureAuth = true

	log.Info().Msgf("Starting IMAP server at localhost:1143")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Msgf("%s", err.Error())
	}
	return nil
}

func main() {
	rootCmd.PersistentFlags().StringP(configurationFileFlag, "f", "", "path to a configuration file")
	err := rootCmd.MarkPersistentFlagRequired(configurationFileFlag)
	if err != nil {
		log.Error().Msgf("failed to set up configuration file flag")
		os.Exit(1)
	}

	rootCmd.Flags().IntP("port", "p", 1143, "port to run imap server on")

	if err = rootCmd.Execute(); err != nil {
		log.Error().Msgf("unable to run application: %s", err.Error())
		if _, e := fmt.Fprintf(os.Stderr, "unable to run application: %s\n", err.Error()); e != nil {
			log.Error().Msgf("unable to write to stderr: %s", e.Error())
			log.Error().Msgf("unable to run application: %s", err.Error())
		}
		os.Exit(1)
	}
	os.Exit(0)
}
