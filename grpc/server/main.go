package main

import (
	"context"
	pb "eztrust/grpc/model"
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type Peer struct {
	PublicKey           string
	Endpoint            string
	AllowedIPs          string
	LatestHandshake     string
	Transfer            string
	PersistentKeepalive string
}

type Interface struct {
	Name          string
	PublicKey     string
	ListeningPort int
	Peers         []Peer
}

type server struct {
	pb.UnimplementedHealthCheckServer
	pb.UnimplementedTunnelInfoServer
}

func parsePort(portStr string) int {
	port, err := strconv.Atoi(portStr)
	if err != nil {
		fmt.Println("Error parsing port:", err)
		return 0
	}
	return port
}

func getPeerStatusRealtime(interfaceName string) (string, error) {
	cmd := exec.Command("wg", "show", interfaceName)
	out, err := cmd.Output()
	if err != nil {
		log.Printf("Failed to get wireguard peer status")
		return "", err
	}
	return string(out), nil
}

func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (server *server) GetTunnelInfo(ctx context.Context, in *pb.GetTunnelRequest) (*pb.GetTunnelResponse, error) {
	log.Printf("Received get tunnel %v info!", in.GetName())
	peerInfoStr, err := getPeerStatusRealtime(in.GetName())
	if err != nil {
		return nil, err
	}

	lines := strings.Split(peerInfoStr, "\n")

	var iface Interface
	var currentPeer *Peer

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "interface:") {
			iface.Name = strings.TrimSpace(strings.TrimPrefix(line, "interface:"))
		} else if strings.HasPrefix(line, "public key:") {
			iface.PublicKey = strings.TrimSpace(strings.TrimPrefix(line, "public key:"))
		} else if strings.HasPrefix(line, "listening port:") {
			fmt.Sscanf(strings.TrimPrefix(line, "listening port:"), "%d", &iface.ListeningPort)
		} else if strings.HasPrefix(line, "peer:") {
			if currentPeer != nil {
				iface.Peers = append(iface.Peers, *currentPeer)
			}
			currentPeer = &Peer{
				PublicKey: strings.TrimSpace(strings.TrimPrefix(line, "peer:")),
			}
		} else if strings.HasPrefix(line, "endpoint:") {
			currentPeer.Endpoint = strings.TrimSpace(strings.TrimPrefix(line, "endpoint:"))
		} else if strings.HasPrefix(line, "allowed ips:") {
			currentPeer.AllowedIPs = strings.TrimSpace(strings.TrimPrefix(line, "allowed ips:"))
		} else if strings.HasPrefix(line, "latest handshake:") {
			currentPeer.LatestHandshake = strings.TrimSpace(strings.TrimPrefix(line, "latest handshake:"))
		} else if strings.HasPrefix(line, "transfer:") {
			currentPeer.Transfer = strings.TrimSpace(strings.TrimPrefix(line, "transfer:"))
		} else if strings.HasPrefix(line, "persistent keepalive:") {
			currentPeer.PersistentKeepalive = strings.TrimSpace(strings.TrimPrefix(line, "persistent keepalive:"))
		}
	}
	if currentPeer != nil {
		iface.Peers = append(iface.Peers, *currentPeer)
	}

	fmt.Printf("%+v\n", iface)

	tunnelResponse := &pb.GetTunnelResponse{}
	tunnelResponse.Name = iface.Name
	tunnelResponse.PublicKey = iface.PublicKey
	tunnelResponse.ListeningPort = fmt.Sprint(iface.ListeningPort)

	for _, peer := range iface.Peers {
		peerInfo := &pb.PeerInfo{
			PublicKey:           peer.PublicKey,
			Endpoint:            peer.Endpoint,
			AllowedIps:          peer.AllowedIPs,
			LatestHandshake:     peer.LatestHandshake,
			Transfer:            peer.Transfer,
			PersistentKeepalive: peer.PersistentKeepalive,
		}
		tunnelResponse.PeerInfo = append(tunnelResponse.PeerInfo, peerInfo)
	}

	return tunnelResponse, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHealthCheckServer(s, &server{})
	pb.RegisterTunnelInfoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
