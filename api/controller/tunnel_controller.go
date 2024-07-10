package controller

import (
	"encoding/json"
	"eztrust/domain"
	"eztrust/infra/queue"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TunnelController struct {
	Rabbitmq *queue.Rabbitmq
}

func (tunnelController *TunnelController) GetTunnelInfo(ctx *gin.Context) {
	tunnelName := ctx.Query("tunnel_name")
	// publish message to rabbitmq
	cmdUuid := uuid.New().String()
	cmd := map[string]interface{}{
		"id":  cmdUuid,
		"cmd": "get_tunnel_info",
		"data": map[string]interface{}{
			"tunnel_name": tunnelName,
		},
	}
	cmdStr, _ := json.Marshal(cmd)
	err := tunnelController.Rabbitmq.Publish("gateway-cmd", []byte(cmdStr), cmdUuid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "Failed to get tunnel info",
			Status:  http.StatusInternalServerError,
		})
		return
	}
	fmt.Println("cmdUuid", cmdUuid)

}
