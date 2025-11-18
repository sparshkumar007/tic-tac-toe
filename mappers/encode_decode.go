package mappers

import (
	service "tic-tac-toe/service/games"

	"github.com/gin-gonic/gin"
)

func DecodeGetGameRequest(c *gin.Context) (service.GetGameRequest, error) {
	var req service.GetGameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return service.GetGameRequest{}, err
	}
	return req, nil
}
