package factory

import "github.com/mateustudeia/payment-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
