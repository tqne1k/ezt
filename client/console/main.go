package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "eztrust/grpc/model"
	"eztrust/internal/wireguard"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ()

var (
	addr        = flag.String("addr", "10.10.0.4:50051", "the address to connect to")
	isLogin     = false
	accessToken = ""
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// Check if user not login, require login
	for {
		if isLogin {
			break
		}
		// Clear screen
		fmt.Print("\033[H\033[2J")
		fmt.Println("----- Login -----")
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		fmt.Println("Password: ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		loginSuccess, respAccessToken := login(username, password)
		if loginSuccess {
			fmt.Println("Login success!")
			// Set access token to global variable
			accessToken = respAccessToken
			isLogin = true
			break
		} else {
			fmt.Println("Login failed. Please try again!")
			time.Sleep(2 * time.Second)
		}
	}

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Println("----- Menu -----")
		fmt.Println("1. Options 1")
		fmt.Println("2. Options 2")
		fmt.Println("3. Exit")
		fmt.Print("Choose the option (1-3): ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			time.Sleep(2 * time.Second)
			fmt.Println("Choose option 1.")
			fmt.Println("Access Token: ", accessToken)
			createTunnel()
			return
		case "2":
			time.Sleep(2 * time.Second)
			fmt.Println("Choose option 2.")
		case "3":
			time.Sleep(2 * time.Second)
			fmt.Println("Exit app.")
			return
		default:
			fmt.Println("Option is not valid!")
			time.Sleep(2 * time.Second)
		}
		fmt.Println()
	}
}

func login(username string, password string) (bool, string) {
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewLoginClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	LoginResponse, err := client.UserLogin(ctx, &pb.LoginRequest{Username: username, Password: password})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}

	if LoginResponse.Status == "true" {
		return true, LoginResponse.AccessToken
	}
	return false, ""
}

func createTunnel() {
	wireguard.CreateInterface("test1")
}
