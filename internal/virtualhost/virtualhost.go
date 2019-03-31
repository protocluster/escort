package vhost

import (
	"context"
	"regexp"
)

var (
	// VirtualHostIDRegex validation expression for VirtualHost IDs
	VirtualHostIDRegex = regexp.MustCompile("^vhost\\/[a-zA-Z0-9\\-\\_]+$")
)

type Rewrite struct {
	Host string
}

type VirtualHost struct {
	ID      string
	Host    string
	Target  string
	Rewrite *Rewrite
}

type VirtualHostRepo interface {
	List(ctx context.Context) ([]VirtualHost, error)
	Get(ctx context.Context, vhostID string) (*VirtualHost, error)
	Delete(ctx context.Context, vhostID string) error
	Update(ctx context.Context, vhost VirtualHost) error
	Revision(ctx context.Context) (int64, error)
}
