// Package credentials produces implementations of a credentials store based on configuration.
package credentials

import (
	"fmt"

	"github.com/shauncampbell/unified-imap/internal/config"
	"github.com/shauncampbell/unified-imap/internal/credentials/memory"
	"github.com/shauncampbell/unified-imap/pkg/credentials"
)

// ForConfiguration loads an authenticator based on the configuration.
func ForConfiguration(cfg *config.Credentials) (credentials.Store, error) {
	if cfg.Type == "in_memory" {
		return memory.New(&cfg.InMemory)
	}
	return nil, fmt.Errorf("no credentials store named '%s'", cfg.Type)
}
