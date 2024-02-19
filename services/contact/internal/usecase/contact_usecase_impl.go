// services/contact/internal/usecase/contact_usecase_impl.go

package usecase

import "github.com/SaraUlan/SaraUlan-03/services/contact/internal/repository"

func NewContactUseCase(repository repository.ContactRepository) ContactUseCase {
    return &ContactUseCaseImpl{Repository: repository}
}
