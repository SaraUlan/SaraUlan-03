package usecase

import (
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/repository"
)

type ContactUseCaseImpl struct {
	Repository repository.ContactRepository
}

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
	return &ContactUseCaseImpl{Repository: repository}
}


func (uc *ContactUseCaseImpl) CreateContact(contact *domain.Contact) (*domain.Contact, error) {

	return contact, nil
}

func (uc *ContactUseCaseImpl) GetContactByID(contactID int) (*domain.Contact, error) {

	return nil, nil
}

func (uc *ContactUseCaseImpl) UpdateContact(contact *domain.Contact) (*domain.Contact, error) {

	return contact, nil
}

func (uc *ContactUseCaseImpl) DeleteContact(contactID int) error {

	return nil
}

func (uc *ContactUseCaseImpl) CreateGroup(group *domain.Group) (*domain.Group, error) {

	return group, nil
}

func (uc *ContactUseCaseImpl) GetGroupByID(groupID int) (*domain.Group, error) {

	return nil, nil
}

func (uc *ContactUseCaseImpl) CreateAndAddContactToGroup(contact *domain.Contact, groupID int) (*domain.Contact, error) {

	return contact, nil
}
