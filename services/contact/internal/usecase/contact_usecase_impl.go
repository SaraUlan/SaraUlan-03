package usecase

import (
	"context"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/repository"
)

type ContactUseCaseImpl struct {
	Repository repository.ContactRepository
}

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
	return &ContactUseCaseImpl{Repository: repository}
}

func (uc *ContactUseCaseImpl) CreateContact(ctx context.Context, contact *domain.Contact) (*domain.Contact, error) {
	return uc.Repository.CreateContact(ctx, contact)
}

func (uc *ContactUseCaseImpl) GetContactByID(ctx context.Context, contactID int) (*domain.Contact, error) {
	return uc.Repository.GetContactByID(ctx, contactID)
}

func (uc *ContactUseCaseImpl) UpdateContact(ctx context.Context, contact *domain.Contact) (*domain.Contact, error) {
	return uc.Repository.UpdateContact(ctx, contact)
}

func (uc *ContactUseCaseImpl) DeleteContact(ctx context.Context, contactID int) error {
	return uc.Repository.DeleteContact(ctx, contactID)
}

func (uc *ContactUseCaseImpl) CreateGroup(ctx context.Context, group *domain.Group) (*domain.Group, error) {
	return uc.Repository.CreateGroup(ctx, group)
}

func (uc *ContactUseCaseImpl) GetGroupByID(ctx context.Context, groupID int) (*domain.Group, error) {
	return uc.Repository.GetGroupByID(ctx, groupID)
}

func (uc *ContactUseCaseImpl) CreateAndAddContactToGroup(ctx context.Context, contact *domain.Contact, groupID int) (*domain.Contact, error) {
	return uc.Repository.CreateAndAddContactToGroup(ctx, contact, groupID)
}
