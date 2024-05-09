package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"

	pb "go-meat-challenges/cmd/meat_counter_api/meatcounter"
)

type server struct {
	pb.UnimplementedMeatCounterServer
}

var fetchTextFromURL = func(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func getMeats() (*pb.GetMeatsResponse, error) {
	text, err := fetchTextFromURL("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return nil, err
	}

	beef := make(map[string]int32)
	for _, word := range strings.Fields(text) {
		beef[word]++
	}

	return &pb.GetMeatsResponse{Beef: beef}, nil
}

func (s *server) GetMeats(ctx context.Context, in *pb.GetMeatsRequest) (*pb.GetMeatsResponse, error) {
	response, err := getMeats()
	if err != nil {
		return nil, err
	}

	return &pb.GetMeatsResponse{Beef: response.Beef}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMeatCounterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
