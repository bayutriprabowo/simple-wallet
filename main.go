package main

import (
	"log"
	"os"

	customerhandler "simple-wallet/customer/handler"
	customerRepo "simple-wallet/customer/repo"
	customerUsecase "simple-wallet/customer/usecase"
	"simple-wallet/setting"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := setting.ConnectDB()

	router := gin.Default()

	customerRepo := customerRepo.CreateCustomerRepo(db)
	customerUsecase := customerUsecase.CreateCustomerUsecase(customerRepo)

	customerhandler.CreateCustomerHandler(router, customerUsecase)
	err := router.Run(":" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
