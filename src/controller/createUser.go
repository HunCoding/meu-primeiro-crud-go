package controller

import (
	"github.com/HunCoding/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Voce chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
