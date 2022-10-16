package main

import (
	"article_GraphQL_gRPC/repo/client"
	"article_GraphQL_gRPC/repo/pb"
	"context"
	"fmt"
	"io"
	"log"
)

// gRPCサーバー動作確認用
func main() {
	c, _ := client.NewClient("localhost:50051")
	list(c)
}

func create(c *client.Client) {
	input := &pb.ArticleInput{
		Author:  "著者1",
		Title:   "タイトル1",
		Content: "内容1",
	}
	res, err := c.Service.CreateArticle(context.Background(), &pb.CreateArticleRequest{ArticleInput: input})
	if err != nil {
		log.Fatalf("Failed to CreateArticle: %v\n", err)
	}
	fmt.Printf("CreateArticles response: %v\n", res)
}

func read(c *client.Client) {
	var id int64 = 1
	res, err := c.Service.ReadArticle(context.Background(), &pb.ReadArticleRequest{Id: id})
	if err != nil {
		log.Fatalf("fatal to read article: %v\n", err)
	}
	fmt.Printf("ReadArticle response: %v\n", res)
}

func update(c *client.Client) {
	var id int64 = 1
	input := &pb.ArticleInput{
		Author:  "著者1new",
		Title:   "タイトル1new",
		Content: "内容1new",
	}
	res, err := c.Service.UpdateArticle(context.Background(), &pb.UpdateArticleRequest{Id: id, ArticleInput: input})
	if err != nil {
		log.Fatalf("fatal to update article: %v\n", err)
	}
	fmt.Printf("updateArticle response: %v\n", res)
}

func delete(c *client.Client) {
	var id int64 = 2
	res, err := c.Service.DeleteArticle(context.Background(), &pb.DeleteArticleRequest{Id: id})
	if err != nil {
		log.Fatalf("failed to delete article: %v\n", err)
	}
	fmt.Printf("DeleteArticle response: %v\n", res)
}

func list(c *client.Client) {
	stream, err := c.Service.ListArticle(context.Background(), &pb.ListArticleRequest{})
	if err != nil {
		log.Fatalf("failed to listArticle: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to server streaming: %v\n", err)
		}
		fmt.Println(res)
	}
}
