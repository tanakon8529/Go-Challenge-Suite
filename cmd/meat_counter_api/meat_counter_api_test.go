package main

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "go-meat-challenges/cmd/meat_counter_api/meatcounter"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterMeatCounterServer(s, &server{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestMeatCounterAPI(t *testing.T) {
	// Mock the fetchTextFromURL function.
	oldFetchTextFromURL := fetchTextFromURL
	defer func() { fetchTextFromURL = oldFetchTextFromURL }()
	fetchTextFromURL = func(url string) (string, error) {
		return "t-bone t-bone t-bone t-bone fatback pastrami", nil
	}

	// Set up the gRPC client.
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewMeatCounterClient(conn)

	// Send a GetMeats request to the server.
	res, err := client.GetMeats(context.Background(), &pb.GetMeatsRequest{})
	if err != nil {
		t.Fatalf("Failed to get meats: %v", err)
	}

	// Check the response.
	expected := map[string]int32{"t-bone": 4, "fatback": 1, "pastrami": 1}
	assert.Equal(t, expected, res.Beef)
}
