package memory

import (
	"time"

	"github.com/emersion/go-imap/backend"
	"github.com/shauncampbell/unified-imap/pkg/source"
)

// Memory is an in memory implementation of the source.Source interface.
type Memory struct {
	username string
}

// ConnectToProvider connects to the specified provider.
func (m *Memory) ConnectToProvider() (source.Provider, error) {
	user := &User{username: m.username}

	body := "From: contact@example.org\r\n" +
		"To: contact@example.org\r\n" +
		"Subject: A little message, just for you\r\n" +
		"Date: Wed, 11 May 2016 14:31:59 +0000\r\n" +
		"Message-ID: <0000000@localhost/>\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hi there :)"

	return &connectedSrc{
		user: &User{
			username: m.username,
			mailboxes: map[string]*Mailbox{
				"INBOX": {
					name: "INBOX",
					user: user,
					Messages: []*Message{
						{
							UID:   6,
							Date:  time.Now(),
							Flags: []string{"\\Seen"},
							Size:  uint32(len(body)),
							Body:  []byte(body),
						},
					},
				},
			},
		},
	}, nil
}

type connectedSrc struct {
	user *User
}

func (s *connectedSrc) GetUser() backend.User {
	return s.user
}

// New returns a new in memory source.
func New(username string, settings map[string]interface{}) *Memory {
	return &Memory{username: username}
}
