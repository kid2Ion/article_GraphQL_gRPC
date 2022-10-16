package client

import (
	"article_GraphQL_gRPC/repo/pb"

	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	Service pb.ArticleServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := pb.NewArticleServiceClient(conn)
	return &Client{conn, c}, nil
}
