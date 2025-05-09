package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mi759/haikuapi/internal/domain"
	"github.com/mi759/haikuapi/internal/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(r *gin.Engine, uc usecase.UserUsecase) {
	h := &UserHandler{usecase: uc}
	r.GET("/user/:id", h.GetUser)
	r.POST("/user", h.CreateUser)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	user, err := h.usecase.GetUserByID(c.Request.Context(), int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var input domain.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := h.usecase.CreateUser(c.Request.Context(), &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, input)
}
