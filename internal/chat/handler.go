package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvtrinh2001/chatapp/proto/chat"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var request chat.CreateRoomRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.Service.CreateRoom(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetRooms(c *gin.Context) {
	request := chat.GetRoomsRequest{}

	response, err := h.Service.GetRooms(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetClients(c *gin.Context) {
	var request chat.GetClientsRequest

	roomId := c.Param("roomId")

	request.RoomId = roomId

	response, err := h.Service.GetClients(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, response)
}
