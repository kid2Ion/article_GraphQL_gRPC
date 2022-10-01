package infra

import (
	"article_GraphQL_gRPC/repo/domain/repository"
	"article_GraphQL_gRPC/repo/pb"
	"context"
	"database/sql"
)

type ArticleRepository struct {
	SqlHandler
}

func NewArticleRepository(sqlHandler SqlHandler) repository.ArticleRepository {
	articleRepository := ArticleRepository{sqlHandler}
	return &articleRepository
}

func (ar *ArticleRepository) Insert(ctx context.Context, input *pb.ArticleInput) (int64, error) {
	cmd := "INSERT INTO articles (author, title, content) VALUES (?, ?, ?)"
	result, err := ar.SqlHandler.Conn.Exec(cmd, input.Author, input.Title, input.Content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ar ArticleRepository) SelectArticleById(ctx context.Context, id int64) (*pb.Article, error) {
	cmd := "SELECT * FROM articles WHERE id = ?"
	row := ar.SqlHandler.Conn.QueryRow(cmd, id)
	// pointaではないかも
	var a *pb.Article

	err := row.Scan(&a.Id, &a.Author, &a.Title, &a.Content)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (ar ArticleRepository) Update(ctx context.Context, id int64, input *pb.ArticleInput) error {
	cmd := "UPDATE articles SET author = ?, title = ?, content = ? WHERE id = ?"
	_, err := ar.SqlHandler.Conn.Exec(cmd, input.Author, input.Title, input.Content, id)
	if err != nil {
		return err
	}

	return nil
}

func (ar ArticleRepository) Delete(ctx context.Context, id int64) error {
	cmd := "DELETE FROM articles WHERE id = ?"
	_, err := ar.SqlHandler.Conn.Exec(cmd, id)
	if err != nil {
		return err
	}

	return nil
}

func (ar ArticleRepository) SelectAllArticles() (*sql.Rows, error) {
	cmd := "SELECT * FROM articles"
	rows, err := ar.SqlHandler.Conn.Query(cmd)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
