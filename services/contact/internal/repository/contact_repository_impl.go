// services/contact/internal/repository/contact_repository_impl.go

package repository

func NewContactRepository() ContactRepository {
    return &ContactRepositoryImpl{}
}
