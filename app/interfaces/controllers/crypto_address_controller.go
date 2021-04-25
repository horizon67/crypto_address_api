package controllers

import (
	"strconv"

	"github.com/horizon67/crypto_address_api/app/domain"
	"github.com/horizon67/crypto_address_api/app/interfaces/database"
	"github.com/horizon67/crypto_address_api/app/usecase"
)

type CryptoAddressController struct {
	Interactor usecase.CryptoAddressInteractor
}

func NewCryptoAddressController(sqlHandler database.SqlHandler) *CryptoAddressController {
	return &CryptoAddressController{
		Interactor: usecase.CryptoAddressInteractor{
			CryptoAddressRepository: &database.CryptoAddressRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CryptoAddressController) Index(c Context) (err error) {
	crypto_addresses, err := controller.Interactor.CryptoAddresses()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, crypto_addresses)
	return
}

func (controller *CryptoAddressController) Show(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	crypto_address, err := controller.Interactor.CryptoAddressById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, crypto_address)
	return
}

func (controller *CryptoAddressController) Create(c Context) (err error) {
	t := domain.CryptoAddress{}
	c.Bind(&t)
	crypto_address, err := controller.Interactor.Add(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, crypto_address)
	return
}

func (controller *CryptoAddressController) Save(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	t, err := controller.Interactor.CryptoAddressById(id)
	c.Bind(&t)
	crypto_address, err := controller.Interactor.Update(t)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, crypto_address)
	return
}

func (controller *CryptoAddressController) Delete(c Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	crypto_address := domain.CryptoAddress{
		ID: id,
	}
	err = controller.Interactor.DeleteById(crypto_address)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, crypto_address)
	return
}
