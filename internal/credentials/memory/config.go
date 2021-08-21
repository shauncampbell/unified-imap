package memory

import "github.com/shauncampbell/unified-imap/pkg/source"

// Config is a configuration struct for the in memory credentials
type Config struct {
	UserSources map[string]source.Config `yaml:"user_sources"`
}
