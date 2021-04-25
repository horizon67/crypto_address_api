package usecase

import "github.com/horizon67/crypto_address_api/app/domain"

type CryptoAddressInteractor struct {
	CryptoAddressRepository CryptoAddressRepository
}

func (interactor *CryptoAddressInteractor) CryptoAddressById(id int) (crypto_address domain.CryptoAddress, err error) {
	crypto_address, err = interactor.CryptoAddressRepository.FindById(id)
	return
}

func (interactor *CryptoAddressInteractor) CryptoAddresses() (crypto_addresses domain.CryptoAddresses, err error) {
	crypto_addresses, err = interactor.CryptoAddressRepository.FindAll()
	return
}

func (interactor *CryptoAddressInteractor) Add(t domain.CryptoAddress) (crypto_address domain.CryptoAddress, err error) {
	crypto_address, err = interactor.CryptoAddressRepository.Store(t)
	return
}

func (interactor *CryptoAddressInteractor) Update(t domain.CryptoAddress) (crypto_address domain.CryptoAddress, err error) {
	crypto_address, err = interactor.CryptoAddressRepository.Update(t)
	return
}

func (interactor *CryptoAddressInteractor) DeleteById(t domain.CryptoAddress) (err error) {
	err = interactor.CryptoAddressRepository.DeleteById(t)
	return
}
