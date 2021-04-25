package usecase

import "github.com/horizon67/crypto_address_api/app/domain"

type CryptoAddressRepository interface {
	FindById(id int) (domain.CryptoAddress, error)
	FindAll() (domain.CryptoAddresses, error)
	Store(domain.CryptoAddress) (domain.CryptoAddress, error)
	Update(domain.CryptoAddress) (domain.CryptoAddress, error)
	DeleteById(domain.CryptoAddress) error
}
