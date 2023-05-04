package routes

import (
	"github.com/brunoOchoa.com/api-REST-FULL/controllers"
	"github.com/gin-gonic/gin"
)

type Route interface {
	RegisterHandlers()
}

type route struct {
	engine     *gin.Engine
	controller controllers.Controller
}

func RegisterHandlers(engine *gin.Engine, controller controllers.Controller) Route {
	return &route{
		engine:     engine,
		controller: controller,
	}
}

func (r route) RegisterHandlers() {
	r.engine.GET("/cliente", r.controller.GetAllClientes())
	r.engine.GET("/cliente/:id", r.controller.GetCliente())
	r.engine.POST("/cliente", r.controller.CreateCliente())
	r.engine.PUT("/cliente/:id", r.controller.UpdateCliente())
	r.engine.DELETE("/cliente/:id", r.controller.DeleteCliente())
}
