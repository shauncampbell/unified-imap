// Package authenticator is a package for looking up authenticator implementations based on configuration.
package authenticator

import (
	"fmt"

	"github.com/shauncampbell/unified-imap/internal/authenticator/memory"
	"github.com/shauncampbell/unified-imap/internal/config"
	"github.com/shauncampbell/unified-imap/pkg/authenticator"
)

// ForConfiguration loads an authenticator based on the configuration.
func ForConfiguration(cfg *config.Authenticator) (authenticator.Authenticator, error) {
	if cfg.Type == "in_memory" {
		return memory.New(&cfg.InMemory), nil
	}
	return nil, fmt.Errorf("no authenticator named '%s'", cfg.Type)
}
