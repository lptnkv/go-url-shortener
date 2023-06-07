package main

import (
	"context"
	"fmt"
	"log"
	"lptnkv/go-url-shortener/proto"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const port = 50051
const shortUrlPrefix = "sh.rt/"

type Record struct {
	originalUrl string
	shortUrl    string
}

type Database interface {
	AddUrl(url string) (string, error)
	GetFullUrl(shortUrl string) (string, error)
}

type server struct {
	proto.UnimplementedShortenerServer
	db Database
}

type postgresDB struct {
	Url string
}

type inMemoryDb struct {
	data    []Record
	counter int
}

func NewPostgresDB() (*postgresDB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("some error occured. Err: %s", err)
	}
	url := os.Getenv("DATABASE_URL")
	return &postgresDB{Url: url}, nil
}

func (db *postgresDB) AddUrl(url string) (shortenedUrl string, err error) {
	conn, err := pgx.Connect(context.Background(), db.Url)
	if err != nil {
		return "", fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Check if url already in db
	row := conn.QueryRow(context.Background(), fmt.Sprintf("select short_url from url where original_url=%s", url))
	var shortUrl string
	row.Scan(&shortUrl)
	if shortUrl != "" {
		return shortUrl, nil
	}

	// Select next generated id to generate hash
	row = conn.QueryRow(context.Background(), "Select currval(pg_get_serial_sequence('url', 'id')) as new_id;")
	var nextId int
	err = row.Scan(&nextId)
	if err != nil {
		return "", fmt.Errorf("unable to get next id from db: %v", err)
	}
	nextId += 1
	hash := idToShortUrl(nextId)
	shortUrl = shortUrlPrefix + hash
	_, err = conn.Exec(context.Background(), fmt.Sprintf("insert into url values (%s, %s);", url, shortUrl))
	if err != nil {
		return "", fmt.Errorf("could not insert url into values: %v", err)
	}
	return shortUrl, nil
}

func (db *postgresDB) GetFullUrl(shortUrl string) (fullUrl string, err error) {
	conn, err := pgx.Connect(context.Background(), db.Url)
	if err != nil {
		return "", fmt.Errorf("unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())
	fullUrlId := shortUrlToId(shortUrl)
	row := conn.QueryRow(context.Background(), fmt.Sprintf("select fullUrl from url where id=%d", fullUrlId))
	var fullUrlFromDb string
	err = row.Scan(&fullUrlFromDb)
	if err != nil {
		return "", fmt.Errorf("unable to get full url from db: %v", err)
	}
	return fullUrlFromDb, nil
}

func idToShortUrl(id int) string {
	alphabet := "abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	base := len(alphabet)
	var res strings.Builder
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

func indexOf(c rune, alphabet string) int {
	for k, v := range alphabet {
		if c == v {
			return k
		}
	}
	return -1
}

func shortUrlToId(shortUrl string) int {
	hash := shortUrl[len(shortUrlPrefix):]
	alphabet := "abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	base := len(alphabet)
	res := 0
	for _, ch := range hash {
		res = (res * base) + indexOf(ch, alphabet)
	}
	return res
}

func (s *server) ShortenUrl(ctx context.Context, req *proto.ShortenUrlRequest) (*proto.ShortenUrlReply, error) {
	url := req.GetUrl()
	shortUrl, err := s.db.AddUrl(url)
	if err != nil {
		return &proto.ShortenUrlReply{Url: shortUrl}, fmt.Errorf("unable to shorten url in grpc request: %v", err)
	}
	return &proto.ShortenUrlReply{Url: shortUrl}, nil
}

func (s *server) GetFullUrl(ctx context.Context, req *proto.GetFullUrlRequest) (*proto.GetFullUrlReply, error) {
	shortUrl := req.GetUrl()
	fullUrl, err := s.db.GetFullUrl(shortUrl)
	if err != nil {
		return &proto.GetFullUrlReply{Url: fullUrl}, fmt.Errorf("unable to get full url in grpc request: %v", err)
	}
	return &proto.GetFullUrlReply{Url: fullUrl}, nil
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
