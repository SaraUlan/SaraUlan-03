// services/contact/internal/delivery/contact_delivery.go

package delivery

import "github.com/SaraUlan/SaraUlan-03/services/contact/internal/domain"

type ContactDelivery interface {
	CreateContact(contact *domain.Contact) (*domain.Contact, error)
	GetContactByID(contactID int) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) (*domain.Contact, error)
	DeleteContact(contactID int) error

	CreateGroup(group *domain.Group) (*domain.Group, error)
	GetGroupByID(groupID int) (*domain.Group, error)
	CreateAndAddContactToGroup(contact *domain.Contact, groupID int) (*domain.Contact, error)
}

