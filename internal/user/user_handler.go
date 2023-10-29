package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var request CreateUserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.Service.CreateUser(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Login(c *gin.Context) {
	var request LoginUserRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userResponse, err := h.Service.Login(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.SetCookie("jwt", userResponse.accessToken, 3600, "/", "localhost", false, true)

	response := &LoginUserResponse{
		ID:       userResponse.ID,
		Username: userResponse.Username,
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logout successfully"})
}
