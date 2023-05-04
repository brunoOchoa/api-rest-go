package controllers

import (
	"net/http"

	"github.com/brunoOchoa.com/api-REST-FULL/requests"
	"github.com/brunoOchoa.com/api-REST-FULL/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Controller interface {
	GetAllClientes() gin.HandlerFunc
	GetCliente() gin.HandlerFunc
	CreateCliente() gin.HandlerFunc
	UpdateCliente() gin.HandlerFunc
	DeleteCliente() gin.HandlerFunc
}

type controller struct {
	service service.ClienteService
}

func NewController(service service.ClienteService) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) GetAllClientes() gin.HandlerFunc {
	return func(context *gin.Context) {
		response, err := c.service.GetAllClientes()
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
		}
		context.JSON(http.StatusOK, response)
	}
}

func (c *controller) GetCliente() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		response, err := c.service.GetCliente(id)

		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
		}
		context.JSON(http.StatusOK, response)
	}
}

func (c *controller) CreateCliente() gin.HandlerFunc {
	return func(context *gin.Context) {
		var createRequest requests.ClienteCreateRequest
		if err := context.ShouldBindJSON(&createRequest); err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		response, err := c.service.CreateCliente(createRequest)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}
func (c *controller) UpdateCliente() gin.HandlerFunc {
	return func(context *gin.Context) {

		var updateRequest requests.ClienteUpdateRequest
		id := context.Param("id")

		if err := context.ShouldBindJSON(&updateRequest); err != nil {
			handleError(context, err, http.StatusBadRequest)
			return
		}
		err := c.service.UpdateCliente(id, updateRequest)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, "Cliente atualizado")
	}
}

func (c *controller) DeleteCliente() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		err := c.service.DeleteCliente(id)
		if err != nil {
			handleError(context, err, http.StatusInternalServerError)
			return
		}
		context.JSON(http.StatusOK, "Cliente deletado")
	}
}

func handleError(context *gin.Context, err error, statusCode int) {
	log.Error(err)
	context.JSON(statusCode, gin.H{
		"error": err.Error(),
	})
}
