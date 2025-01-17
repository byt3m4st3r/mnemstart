package auth

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
)

const (
	SessionName = "session"
)

type SessionOptions struct {
	StorePath  string
	CookiesKey string
	MaxAge     int
	HttpOnly   bool // Should be true if the site is served over HTTP (development environment)
	Secure     bool // Should be true if the site is served over HTTPS (production environment)
}

func NewFileStore(opts SessionOptions) (*sessions.FilesystemStore, error) {
	if err := os.MkdirAll(opts.StorePath, 0700); err != nil {
		return nil, fmt.Errorf("failed to create session storage directory: %w", err)
	}

	store := sessions.NewFilesystemStore(opts.StorePath, []byte(opts.CookiesKey))

	store.MaxAge(opts.MaxAge)
	store.MaxLength(8192)
	store.Options.Path = "/"
	store.Options.HttpOnly = opts.HttpOnly
	store.Options.Secure = opts.Secure

	return store, nil
}
