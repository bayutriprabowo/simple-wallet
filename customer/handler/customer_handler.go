package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"simple-wallet/constant"
	"simple-wallet/customer"
	"simple-wallet/model"
	"simple-wallet/utils"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	customerUsecase customer.CustomerUsecase
}

func CreateCustomerHandler(r *gin.Engine, customerUsecase customer.CustomerUsecase) {
	customerHandler := CustomerHandler{customerUsecase}

	r.POST("/account", customerHandler.addCustomer)
	r.GET("/account", customerHandler.GetAccounts)
	r.GET("/account/:"+constant.AccountNumber, customerHandler.getAccountDetail)
	r.PUT("/transfer", customerHandler.transferBalance)
}

func (e *CustomerHandler) addCustomer(c *gin.Context) {
	var account model.Account
	// err := c.Bind(&account)
	// if err != nil {
	// 	fmt.Printf("[CustomerHandler.addCustomer] error bind data %v \n", err)
	// 	utils.HandleError(c, http.StatusInternalServerError, constant.ServerHasWrong)
	// 	return
	// }

	b, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(b, &account)
	if account.CustomerName == "" || account.Balance == 0 {
		utils.HandleError(c, http.StatusBadRequest, constant.FieldsAreRequired)
		return
	}
	isSuccess := e.customerUsecase.InsertCustomer(&account)
	if !isSuccess {
		utils.HandleError(c, http.StatusInternalServerError, constant.ServerHasWrong)
		return
	}
	utils.HandleSuccess(c, constant.SuccessInputData)
}

func (e *CustomerHandler) getAccountDetail(c *gin.Context) {
	accountNumber := c.Param(constant.AccountNumber)
	isExist := e.customerUsecase.CheckAccoutExist(accountNumber)
	if !isExist {
		utils.HandleError(c, http.StatusNotFound, constant.AccountNumberNotFound)
		return
	}
	account, err := e.customerUsecase.FindCustomerByID(accountNumber)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccessData(c, account)
}

func (e *CustomerHandler) GetAccounts(c *gin.Context) {
	customers, err := e.customerUsecase.FindCustomers()
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSuccessData(c, customers)
}

func (e *CustomerHandler) transferBalance(c *gin.Context) {
	myAccountNumber := c.GetHeader(constant.AccountNumber)
	var account model.Transfer
	err := c.Bind(&account)
	if err != nil {
		fmt.Printf("[CustomerHandler.TransferBalance] error bind data %v \n", err)
		utils.HandleError(c, http.StatusInternalServerError, constant.ServerHasWrong)
		return
	}
	isExist := e.customerUsecase.CheckAccoutExist(myAccountNumber)
	if !isExist {
		utils.HandleError(c, http.StatusNotFound, constant.AccountNumberNotFound)
		return
	}
	isValid := e.customerUsecase.CheckAccoutExist(account.ToAccountNumber)
	if !isValid {
		utils.HandleError(c, http.StatusNotFound, constant.AccountNumberTrasferNotFound)
		return
	}
	account.MyAccountNumber = myAccountNumber
	isSucces := e.customerUsecase.TransferBalance(&account)
	if !isSucces {
		utils.HandleError(c, http.StatusBadRequest, constant.InsufficientBalance)
		return
	}
	utils.HandleSuccess(c, constant.SuccessTransferAmount)
}
