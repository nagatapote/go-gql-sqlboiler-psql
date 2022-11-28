package router

import (
	"go-gql-sqlboiler-psql/infrastructure/db"
	"go-gql-sqlboiler-psql/infrastructure/graphql"
	"go-gql-sqlboiler-psql/interface/database"
	"go-gql-sqlboiler-psql/interface/resolvers"
	"go-gql-sqlboiler-psql/usecase/converter"
	"go-gql-sqlboiler-psql/usecase/godevusecase"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

func HandlerPlayGround() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func HandlerGraphql(e *gin.Engine) {
	root := e.Group("/query")

	dbUtil := db.NewDbUtil(db.DB)
	appRepositories := database.NewAppRepositories(dbUtil)
	appConverters := converter.NewAppConverters()

	c := graphql.Config{}

	c.Resolvers = CreateResolver(appRepositories, appConverters)

	h := handler.NewDefaultServer(graphql.NewExecutableSchema(c))

	root.POST(
		"",
		func(c *gin.Context) {
			h.ServeHTTP(c.Writer, c.Request)
		},
	)
}

func CreateResolver(repositories *database.AppRepositories, converters *converter.AppConverters) *resolvers.Resolver {
	userUseCase := godevusecase.NewUserUseCase(repositories.UserRepository, converters.UserConverter)

	return &resolvers.Resolver{
		UserUseCase:                userUseCase,
	}
}