
package repository

import (
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
	"context"
)
type ContactRepository interface {
	CreateContact(context.Context, *domain.Contact) (*domain.Contact, error)
	GetContactByID(context.Context, int) (*domain.Contact, error)
	UpdateContact(context.Context, *domain.Contact) (*domain.Contact, error)
	DeleteContact(context.Context, int) error

	CreateGroup(context.Context, *domain.Group) (*domain.Group, error)

	GetGroupByID(context.Context, int) (*domain.Group, error)

 	AddContactToGroup(context.Context, int, int) error
    CreateAndAddContactToGroup(ctx context.Context, contact *domain.Contact, groupID int) (*domain.Contact, error)

}
