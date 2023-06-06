package main

import (
	"context"
	"fmt"
	"log"
	"lptnkv/go-url-shortener/proto"
	"net"
	"strings"

	"google.golang.org/grpc"
)

const port = 50051

var counter int

type record struct {
	originalUrl string
	shortUrl    string
}

type server struct {
	proto.UnimplementedShortenerServer
}

func idToShortUrl(id int) string {
	alphabet := "abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	base := len(alphabet)
	var res strings.Builder
	fmt.Println(base)
	var digits []int
	for id > 0 {
		remainder := id % base
		digits = append(digits, remainder)
		id = id / base
	}
	for i := len(digits) - 1; i >= 0; i-- {
		char := alphabet[digits[i]]
		res.WriteByte(char)
	}
	return res.String()
}

func (s *server) ShortenUrl(ctx context.Context, req *proto.ShortenUrlRequest) (*proto.ShortenUrlReply, error) {
	url := "sh.rt/qweqw"
	return &proto.ShortenUrlReply{Url: url}, nil
}

func (s *server) GetFullUrl(ctx context.Context, req *proto.GetFullUrlRequest) (*proto.GetFullUrlReply, error) {
	url := "sh.rt/qweqw"
	return &proto.GetFullUrlReply{Url: url}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterShortenerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
