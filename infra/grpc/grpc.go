package grpc

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"eztrust/domain"
	pb "eztrust/grpc/model"
	"eztrust/infra/keycloak"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "10.10.0.1:50051", "the address to connect to")
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedLoginServer
}

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

func (server *server) UserLogin(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received login request: %v", in.GetUsername())
	loginSuccess, accessToken := keycloak.ValidateUser(in.GetUsername(), in.GetPassword())
	if !loginSuccess {
		return &pb.LoginResponse{Status: "false"}, nil
	}
	return &pb.LoginResponse{Status: "true", AccessToken: accessToken}, nil
}

func RunServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
