package controller

import (
	"net/http"

	"enigmacamp.com/golang-sample/model"
	"enigmacamp.com/golang-sample/usecase"
	"enigmacamp.com/golang-sample/utils"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router  *gin.Engine
	usecase usecase.CustomerUsecase
}

func (cc *CustomerController) getAllCustomer(ctx *gin.Context) {
	customers, err := cc.usecase.GetAllCustomer()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}
func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	customers, err := cc.usecase.FindCustomerById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) registerCustomer(ctx *gin.Context) {
	var customer model.Customer
	customer.Id = utils.UuidGenerate()
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := cc.usecase.RegisterCustomer(customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func NewCustomerController(r *gin.Engine, usecase usecase.CustomerUsecase) *CustomerController {
	controller := CustomerController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/customer", controller.getAllCustomer)
	r.GET("/customer/:id", controller.getCustomerById)
	r.POST("/customer", controller.registerCustomer)
	return &controller
}
