package repository

import (
	"article_GraphQL_gRPC/repo/pb"
	"context"
	"database/sql"
)

type ArticleRepository interface {
	Insert(ctx context.Context, input *pb.ArticleInput) (int64, error)
	SelectArticleById(ctx context.Context, id int64) (*pb.Article, error)
	Update(ctx context.Context, id int64, input *pb.ArticleInput) error
	Delete(ctx context.Context, id int64) error
	SelectAllArticles() (*sql.Rows, error)
}
