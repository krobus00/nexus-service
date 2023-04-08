package bootstrap

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-jwt/jwt/v4"
	authPB "github.com/krobus00/auth-service/pb/auth"
	"github.com/krobus00/nexus-service/internal/config"
	"github.com/krobus00/nexus-service/internal/constant"
	"github.com/krobus00/nexus-service/internal/graph"
	"github.com/krobus00/nexus-service/internal/graph/directive"
	"github.com/krobus00/nexus-service/internal/infrastructure"
	"github.com/krobus00/nexus-service/internal/model"
	productPB "github.com/krobus00/product-service/pb/product"
	storagePB "github.com/krobus00/storage-service/pb/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartServer() {
	e := infrastructure.NewEcho()

	// init grpc client
	authConn, err := grpc.Dial(config.AuthGRPCHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	continueOrFatal(err)
	authClient := authPB.NewAuthServiceClient(authConn)

	storageConn, err := grpc.Dial(config.StorageGRPCHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	continueOrFatal(err)
	storageClient := storagePB.NewStorageServiceClient(storageConn)

	productConn, err := grpc.Dial(config.ProductGRPCHost(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	continueOrFatal(err)
	productClient := productPB.NewProductServiceClient(productConn)

	// init resolver
	resolver := graph.NewResolver()
	err = resolver.InjectAuthClient(authClient)
	continueOrFatal(err)
	err = resolver.InjectStorageClient(storageClient)
	continueOrFatal(err)
	err = resolver.InjectProductClient(productClient)
	continueOrFatal(err)

	c := graph.Config{Resolvers: resolver}
	c.Directives.Binding = directive.Binding

	graphQLHandler := handler.NewDefaultServer(graph.NewExecutableSchema(c))
	playgroundHandler := playground.Handler("GraphQL", "/query")
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
	}))
	e.Use(decodeJWTToken())
	e.POST("/query", func(c echo.Context) error {
		graphQLHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	if config.Env() == "development" {
		log.Info("register playground endpoint")
		e.GET("/playground", func(c echo.Context) error {
			playgroundHandler.ServeHTTP(c.Response(), c.Request())
			return nil
		})
	}

	go func() {
		_ = e.Start(":" + config.HTTPPort())
	}()
	log.Info(fmt.Sprintf("http server started on :%s", config.HTTPPort()))

	wait := gracefulShutdown(context.Background(), config.GracefulShutdownTimeOut(), map[string]operation{
		"http": func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})
	<-wait
}

func decodeJWTToken() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			res := model.NewResponse().WithMessage(model.ErrTokenInvalid.Error())
			accessToken := eCtx.Request().Header.Get("Authorization")
			accessToken = strings.ReplaceAll(accessToken, "Bearer ", "")

			token, _ := jwt.Parse(accessToken, nil)
			if token == nil {
				return next(eCtx)
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}

			userID, ok := claims["userID"]
			if !ok {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}

			tokenID, ok := claims["jti"]
			if !ok {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}

			ctx := context.WithValue(eCtx.Request().Context(), constant.KeyUserIDCtx, userID)
			ctx = context.WithValue(ctx, constant.KeyTokenIDCtx, tokenID)
			eCtx.SetRequest(eCtx.Request().WithContext(ctx))
			return next(eCtx)
		}
	}
}
