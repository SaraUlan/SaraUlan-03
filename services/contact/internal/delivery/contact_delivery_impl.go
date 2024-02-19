// services/contact/internal/delivery/contact_delivery_impl.go

package delivery

import "github.com/SaraUlan/SaraUlan-03/services/contact/internal/usecase"

func NewContactDelivery(useCase usecase.ContactUseCase) ContactDelivery {
    return &ContactDeliveryImpl{UseCase: useCase}
}
