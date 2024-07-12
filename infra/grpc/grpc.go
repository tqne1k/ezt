package grpc

import (
	"context"
	"flag"
	"log"
	"time"

	"eztrust/domain"
	pb "eztrust/grpc/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "10.10.0.1:50051", "the address to connect to")
)

func GetTunnelInfo(tunnelName string) (domain.NetworkResponse, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tunnelClient := pb.NewTunnelInfoClient(conn)

	tunnelInfo, err := tunnelClient.GetTunnelInfo(ctx, &pb.GetTunnelRequest{Name: "wg0"})
	if err != nil {
		log.Fatalf("could not get tunnel info: %v", err)
	}

	network := domain.NetworkResponse{
		Name:          tunnelInfo.GetName(),
		PublicKey:     tunnelInfo.GetPublicKey(),
		ListeningPort: tunnelInfo.GetListeningPort(),
	}

	for _, peer := range tunnelInfo.GetPeerInfo() {
		network.Devices = append(network.Devices, domain.Device{
			PublicKey:       peer.GetPublicKey(),
			Endpoint:        peer.GetEndpoint(),
			AllowedIPs:      peer.GetAllowedIps(),
			LatestHandshake: peer.GetLatestHandshake(),
			Transfer:        peer.GetTransfer(),
		})
	}

	return network, nil
}
