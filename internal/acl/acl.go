package acl

import (
	"context"
	"regexp"
)

var (
	// ACLIDRegex validation expression for ACL IDs
	ACLIDRegex = regexp.MustCompile("^acl\\/[a-zA-Z0-9\\-\\_]+$")
)

type ACL struct {
	ID            string
	VirtualHostID string
	PathRegexp    string
	SubjectIDs    []string
}

type ACLRepo interface {
	List(ctx context.Context) ([]ACL, error)
	Get(ctx context.Context, aclID string) (*ACL, error)
	Delete(ctx context.Context, aclID string) error
	Update(ctx context.Context, acl ACL) error
	Revision(ctx context.Context) (int64, error)
}
