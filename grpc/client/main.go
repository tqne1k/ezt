package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "eztrust/grpc/model"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "10.10.0.1:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHealthCheckClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHi(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	tunnelClient := pb.NewTunnelInfoClient(conn)

	tunnelInfo, err := tunnelClient.GetTunnelInfo(ctx, &pb.GetTunnelRequest{Name: "wg0"})
	if err != nil {
		log.Fatalf("could not get tunnel info: %v", err)
	}
	log.Printf("Tunnel info: %v", tunnelInfo)
}
