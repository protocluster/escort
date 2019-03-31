package group

import (
	"context"
	"regexp"
)

var (
	// GroupIDRegex validation expression for Group IDs
	GroupIDRegex = regexp.MustCompile("^group\\/[a-zA-Z0-9\\-\\_]+$")

	// UserIDRegex validation expression for User IDs/email addresses
	UserIDRegex = regexp.MustCompile("^user\\/.+@.+$")
)

type Group struct {
	ID         string
	SubjectIDs []string
}

type GroupRepo interface {
	List(ctx context.Context) ([]Group, error)
	Get(ctx context.Context, groupID string) (*Group, error)
	Delete(ctx context.Context, groupID string) error
	Update(ctx context.Context, group Group) error
	Revision(ctx context.Context) (int64, error)
}
