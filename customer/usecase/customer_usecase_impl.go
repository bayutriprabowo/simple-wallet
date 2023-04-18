package usecase

import (
	"simple-wallet/customer"
	"simple-wallet/model"
)

type CustomerUsecaseImpl struct {
	customerRepo customer.CustomerRepo
}

func CreateCustomerUsecase(customerRepo customer.CustomerRepo) customer.CustomerUsecase {
	return &CustomerUsecaseImpl{customerRepo}
}

func (e *CustomerUsecaseImpl) FindCustomerByID(accountNumber string) (*model.Account, error) {
	return e.customerRepo.FindCustomerByID(accountNumber)
}

func (e *CustomerUsecaseImpl) InsertCustomer(customer *model.Account) bool {
	return e.customerRepo.InsertCustomer(customer)
}

func (e *CustomerUsecaseImpl) TransferBalance(balance *model.Transfer) bool {
	return e.customerRepo.TransferBalance(balance)
}

func (e *CustomerUsecaseImpl) CheckAccoutExist(accountNumber string) bool {
	return e.customerRepo.CheckAccoutExist(accountNumber)
}

func (e *CustomerUsecaseImpl) CheckCustomerExist(customerNumber string) bool {
	return e.customerRepo.CheckCustomerExist(customerNumber)
}

func (e *CustomerUsecaseImpl) FindCustomers() (*[]model.Account, error) {
	return e.customerRepo.FindCustomers()
}
