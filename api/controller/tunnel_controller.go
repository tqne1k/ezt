package controller

import (
	"eztrust/infra/grpc"
	"eztrust/infra/queue"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TunnelController struct {
	Rabbitmq *queue.Rabbitmq
}

func (tunnelController *TunnelController) GetTunnelInfo(ctx *gin.Context) {
	tunnelName := ctx.Query("tunnel_name")
	// Get tunnel info from grpc server
	network, err := grpc.GetTunnelInfo(tunnelName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": network})

}
