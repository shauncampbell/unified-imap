// Package source provides interfaces for working with different mailbox sources.
package source

import "github.com/emersion/go-imap/backend"

// Source is an interface which allows connections to a mailbox
type Source interface {
	ConnectToProvider() (Provider, error)
}

// Provider is an interface which represents an email provider.
type Provider interface {
	GetUser() backend.User
}
