package injector

import (
	"article_GraphQL_gRPC/repo/domain/repository"
	"article_GraphQL_gRPC/repo/infra"
	"article_GraphQL_gRPC/repo/service"
)

func injectDB() infra.SqlHandler {
	sqlHandler := infra.NewSqlHandler()
	return *sqlHandler
}

func injectArticleRepository() repository.ArticleRepository {
	sqlHandler := injectDB()
	return infra.NewArticleRepository(sqlHandler)
}

func InjectArticleService() service.ArticleService {
	repository := injectArticleRepository()
	return service.NewArticleServer(repository)
}
