package main

import (
	"context"
	"log"
	"lptnkv/go-url-shortener/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewShortenerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ShortenUrl(ctx, &proto.ShortenUrlRequest{})
	if err != nil {
		log.Fatalf("could not shorten url: %v", err)
	}
	log.Printf("Shortened url: %s", r.GetUrl())
}
