package delivery

import (
	"context"
	"errors"
	"log"
	"github.com/google/uuid"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/usecase"
)

type ContactDeliveryImpl struct {
	UseCase usecase.ContactUseCase
}
var ErrContactNotFound = errors.New("Contact not found")


func NewContactDelivery(useCase usecase.ContactUseCase) ContactDelivery {
	return &ContactDeliveryImpl{UseCase: useCase}
}

func (cd *ContactDeliveryImpl) CreateContact(contact *domain.Contact) (*domain.Contact, error) {
	ctx := context.WithValue(context.Background(), "ID", uuid.New().String())
	return cd.UseCase.CreateContact(ctx, contact)
}

func (cd *ContactDeliveryImpl) GetContactByID(contactID int) (*domain.Contact, error) {
	ctx := context.Background()

	existingContact, err := cd.UseCase.GetContactByID(ctx, contactID)
	if err != nil {
		log.Printf("Ошибка при получении контакта по ID: %v", err)
		return nil, err
	}

	if existingContact == nil {
		return nil, ErrContactNotFound
	}

	return existingContact, nil
}

func (cd *ContactDeliveryImpl) UpdateContact(contact *domain.Contact) (*domain.Contact, error) {
	ctx := context.Background()
	return cd.UseCase.UpdateContact(ctx, contact)
}

func (cd *ContactDeliveryImpl) DeleteContact(contactID int) error {
	ctx := context.Background()
	return cd.UseCase.DeleteContact(ctx, contactID)
}

func (cd *ContactDeliveryImpl) CreateGroup(group *domain.Group) (*domain.Group, error) {
	ctx := context.WithValue(context.Background(), "ID", uuid.New().String())
	return cd.UseCase.CreateGroup(ctx, group)
}

func (cd *ContactDeliveryImpl) GetGroupByID(groupID int) (*domain.Group, error) {
	ctx := context.Background()
	return cd.UseCase.GetGroupByID(ctx, groupID)
}

func (cd *ContactDeliveryImpl) CreateAndAddContactToGroup(contact *domain.Contact, groupID int) (*domain.Contact, error) {
	ctx := context.Background()
	return cd.UseCase.CreateAndAddContactToGroup(ctx, contact, groupID)
}
