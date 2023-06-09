package customer

import "simple-wallet/model"

type CustomerUsecase interface {
	FindCustomerByID(accountNumber string) (*model.Account, error)
	InsertCustomer(customer *model.Account) bool
	CheckAccoutExist(accountNumber string) bool
	CheckCustomerExist(customerNumber string) bool
	TransferBalance(balance *model.Transfer) bool
	FindCustomers() (*[]model.Account, error)
}
