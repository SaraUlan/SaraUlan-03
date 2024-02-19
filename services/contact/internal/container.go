// services/contact/internal/container.go

package internal

import (
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/delivery"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/repository"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/usecase"
)

type Container struct {
	ContactRepository repository.ContactRepository
	ContactUseCase    usecase.ContactUseCase
	ContactDelivery   delivery.ContactDelivery
}

func NewContainer() *Container {
	contactRepository := repository.NewContactRepository()

	contactUseCase := usecase.NewContactUseCase(contactRepository)

	contactDelivery := delivery.NewContactDelivery(contactUseCase)

	return &Container{
		ContactRepository: contactRepository,
		ContactUseCase:    contactUseCase,
		ContactDelivery:   contactDelivery,
	}
}
