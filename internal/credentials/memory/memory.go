// Package memory provides an in memory credentials store.
package memory

import (
	"fmt"

	unifiedSource "github.com/shauncampbell/unified-imap/internal/source"
	"github.com/shauncampbell/unified-imap/pkg/credentials"

	"github.com/emersion/go-imap/backend"
	"github.com/shauncampbell/unified-imap/pkg/source"
)

// Store is an in memory credentials store.
type Store struct {
	users map[string]source.Source
}

// GetBackendForUser returns a backend user session for the specified username.
func (s *Store) GetBackendForUser(username string) (backend.User, error) {
	if s.users[username] != nil {
		src, err := s.users[username].ConnectToProvider()
		if err != nil {
			return nil, fmt.Errorf("failed to connect to source: %w", err)
		}

		return src.GetUser(), nil
	}
	return nil, fmt.Errorf("no backends defined for user")
}

// New provides a new in memory credentials store.
func New(config *Config) (credentials.Store, error) {
	users := make(map[string]source.Source)
	for username, srcCfg := range config.UserSources {
		cfg := srcCfg
		src, err := unifiedSource.ForConfiguration(username, &cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to load source for user '%s': %s", username, err.Error())
		}
		users[username] = src
	}
	return &Store{users: users}, nil
}
