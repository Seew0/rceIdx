package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seew0/rceIdx/domain/models"
	"github.com/seew0/rceIdx/internal/logic"
)

type Handler struct {
	logic logic.Logic
}

func NewHandler(logic logic.Logic) *Handler {
	return &Handler{logic: logic}
}

func (h *Handler) GetHealth(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func (h *Handler) Register(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error while binding json to object": err.Error(),
		})
		return
	}

	err = h.logic.InsertUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error while inserting data": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "added user successfully",
	})
}
