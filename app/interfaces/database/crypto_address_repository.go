package database

import "github.com/horizon67/crypto_address_api/app/domain"

type CryptoAddressRepository struct {
	SqlHandler
}

func (repo *CryptoAddressRepository) FindById(id int) (crypto_address domain.CryptoAddress, err error) {
	if err = repo.Find(&crypto_address, id).Error; err != nil {
		return
	}
	return
}

func (repo *CryptoAddressRepository) FindAll() (crypto_addresses domain.CryptoAddresses, err error) {
	if err = repo.Find(&crypto_addresses).Error; err != nil {
		return
	}
	return
}

func (repo *CryptoAddressRepository) Store(t domain.CryptoAddress) (crypto_address domain.CryptoAddress, err error) {
	if err = repo.Create(&t).Error; err != nil {
		return
	}
	crypto_address = t
	return
}

func (repo *CryptoAddressRepository) Update(t domain.CryptoAddress) (crypto_address domain.CryptoAddress, err error) {
	if err = repo.Save(&t).Error; err != nil {
		return
	}
	crypto_address = t
	return
}

func (repo *CryptoAddressRepository) DeleteById(crypto_address domain.CryptoAddress) (err error) {
	if err = repo.Delete(&crypto_address).Error; err != nil {
		return
	}
	return
}
