package internal

import (
	"database/sql"

	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/delivery"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/repository"
	"github.com/SaraUlan/SaraUlan-03/services/contact/internal/usecase"
)

type Container struct {
	ContactRepository repository.ContactRepository
	ContactUseCase    usecase.ContactUseCase
	ContactDelivery   delivery.ContactDelivery
}

func NewContainer(db *sql.DB) *Container {
	contactRepository := repository.NewContactRepository(db)

	contactUseCase := usecase.NewContactUseCase(contactRepository)

	contactDelivery := delivery.NewContactDelivery(contactUseCase)

	return &Container{
		ContactRepository: contactRepository,
		ContactUseCase:    contactUseCase,
		ContactDelivery:   contactDelivery,
	}
}
