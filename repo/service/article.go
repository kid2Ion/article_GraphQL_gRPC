package service

import (
	"article_GraphQL_gRPC/repo/domain/repository"
	"article_GraphQL_gRPC/repo/pb"
	"context"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error)
	ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error)
	UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error)
	DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error)
	ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error
}

type articleService struct {
	repository repository.ArticleRepository
}

func NewArticleServer(r repository.ArticleRepository) ArticleService {
	return &articleService{r}
}

func (as *articleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
	input := req.GetArticleInput()

	id, err := as.repository.Insert(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (as articleService) ReadArticle(ctx context.Context, req *pb.ReadArticleRequest) (*pb.ReadArticleResponse, error) {
	id := req.GetId()
	article, err := as.repository.SelectArticleById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.ReadArticleResponse{
		Article: &pb.Article{
			Id:      article.Id,
			Author:  article.Author,
			Title:   article.Title,
			Content: article.Content,
		},
	}, nil
}

func (as articleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
	id := req.GetId()
	input := req.GetArticleInput()

	if err := as.repository.Update(ctx, id, input); err != nil {
		return nil, err
	}

	return &pb.UpdateArticleResponse{
		Article: &pb.Article{
			Id:      id,
			Author:  input.Author,
			Title:   input.Title,
			Content: input.Content,
		},
	}, nil
}

func (as articleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
	id := req.GetId()

	if err := as.repository.Delete(ctx, id); err != nil {
		return nil, err
	}

	return &pb.DeleteArticleResponse{Id: id}, nil
}

func (as articleService) ListArticle(req *pb.ListArticleRequest, stream pb.ArticleService_ListArticleServer) error {
	rows, err := as.repository.SelectAllArticles()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var a pb.Article
		err := rows.Scan(&a.Id, &a.Author, &a.Title, &a.Content)
		if err != nil {
			return err
		}

		stream.Send(&pb.ListArticleResponse{Article: &a})
	}
	return nil
}
