package usecase

import (
    "context"
    "github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
)

type ContactUseCase interface {
	CreateContact(ctx context.Context, contact *domain.Contact) (*domain.Contact, error)
	GetContactByID(ctx context.Context, contactID int) (*domain.Contact, error)
	UpdateContact(ctx context.Context, contact *domain.Contact) (*domain.Contact, error)
	DeleteContact(ctx context.Context, contactID int) error

	CreateGroup(ctx context.Context, group *domain.Group) (*domain.Group, error)
	GetGroupByID(ctx context.Context, groupID int) (*domain.Group, error)

	CreateAndAddContactToGroup(ctx context.Context, contact *domain.Contact, groupID int) (*domain.Contact, error)
}
