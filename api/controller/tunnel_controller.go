package controller

import (
	"eztrust/domain"
	"eztrust/infra/grpc"
	"eztrust/infra/queue"
	"eztrust/internal/wireguard"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TunnelController struct {
	Rabbitmq *queue.Rabbitmq
	Database *gorm.DB
}

func (tunnelController *TunnelController) GetTunnelInfo(ctx *gin.Context) {
	tunnelName := ctx.Query("tunnel_name")
	// Get tunnel info from grpc server
	NetworkResponse, err := grpc.GetTunnelInfo(tunnelName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map network data
	var network domain.Network
	queryResult := tunnelController.Database.Where("name = ?", tunnelName).First(&network)
	if queryResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": queryResult.Error.Error()})
		return
	}
	NetworkResponse.NetworkAddress = network.NetworkAddress

	// Map device name
	for i, peer := range NetworkResponse.Devices {
		// Find device by public key
		var device domain.Device
		queryResult := tunnelController.Database.Where("public_key = ?", peer.PublicKey).First(&device)
		if queryResult.Error != nil {
			continue
		}
		NetworkResponse.Devices[i].PublicKey = device.Name
	}

	ctx.HTML(http.StatusOK, "tunnel-management.tmpl", gin.H{
		"title":   "Tunnel Management",
		"network": NetworkResponse,
	})
}

func (tunnelController *TunnelController) AddDevice(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "add-device.tmpl", gin.H{
		"title": "Tunnel Management",
	})
}

func (tunnelController *TunnelController) CreateDevice(ctx *gin.Context) {
	deviceName := ctx.PostForm("device_name")

	// Generate key pair
	privateKey := wireguard.GeneratePrivateKey()
	publicKey := wireguard.GeneratePublicKey(privateKey)

	device := domain.Device{
		Id:         uuid.New().String(),
		Name:       deviceName,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}

	// Save device to database
	result := tunnelController.Database.Create(&device)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"device_name": deviceName,
	})
}
