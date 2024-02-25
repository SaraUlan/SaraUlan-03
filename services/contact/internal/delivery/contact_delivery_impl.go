package delivery

import (
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/usecase"
)

type ContactDeliveryImpl struct {
	UseCase usecase.ContactUseCase
}

func NewContactDelivery(useCase usecase.ContactUseCase) ContactDelivery {
	return &ContactDeliveryImpl{UseCase: useCase}
}

func (cd *ContactDeliveryImpl) CreateContact(contact *domain.Contact) (*domain.Contact, error) {
	return cd.UseCase.CreateContact(contact)
}

func (cd *ContactDeliveryImpl) GetContactByID(contactID int) (*domain.Contact, error) {
	return cd.UseCase.GetContactByID(contactID)
}

func (cd *ContactDeliveryImpl) UpdateContact(contact *domain.Contact) (*domain.Contact, error) {
	return cd.UseCase.UpdateContact(contact)
}

func (cd *ContactDeliveryImpl) DeleteContact(contactID int) error {
	return cd.UseCase.DeleteContact(contactID)
}

func (cd *ContactDeliveryImpl) CreateGroup(group *domain.Group) (*domain.Group, error) {
	return cd.UseCase.CreateGroup(group)
}

func (cd *ContactDeliveryImpl) GetGroupByID(groupID int) (*domain.Group, error) {
	return cd.UseCase.GetGroupByID(groupID)
}

func (cd *ContactDeliveryImpl) CreateAndAddContactToGroup(contact *domain.Contact, groupID int) (*domain.Contact, error) {
	return cd.UseCase.CreateAndAddContactToGroup(contact, groupID)
}
