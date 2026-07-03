package handler

import (
	"net/http"

	"github.com/buibahoanvu/ebvn/internal/service"
	"github.com/gin-gonic/gin"
)


const passwordlength = 12

type GenPass interface {
	GeneratePassword (c *gin.Context)
}
type genPassHandler struct {
	genPassService service.GenPass
}

func NewGenPass(genPassSvc service.GenPass) GenPass {
	return &genPassHandler{
		genPassService: genPassSvc,
	}
}

func (s *genPassHandler) GeneratePassword (c *gin.Context){
	pass , err := s.genPassService.GeneratePassword(passwordlength)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Err"})
	}

	c.JSON(http.StatusOK, gin.H{"password": pass})
}