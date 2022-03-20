package factory

import "github.com/mateustudeia/gateway-pagamento/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
