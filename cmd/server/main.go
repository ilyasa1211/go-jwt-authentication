package main

import (
	"net/http"

	"github.com/ilyasa1211/go-jwt-authentication/internal/infrastructure/database/sqlite"
	"github.com/ilyasa1211/go-jwt-authentication/internal/infrastructure/database/sqlite/repositories"
	handler "github.com/ilyasa1211/go-jwt-authentication/internal/infrastructure/http"
	"github.com/ilyasa1211/go-jwt-authentication/internal/middlewares"
	"github.com/ilyasa1211/go-jwt-authentication/internal/services"
)

func main() {
	db := sqlite.NewSqliteConn("data/data.db")

	userRepo := repositories.NewUserRepository(db)
	userSvc := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSvc)

	authSvc := services.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authSvc)

	http.HandleFunc("GET /users", middlewares.ChainMiddlewares(userHandler.Index, middlewares.AuthMiddleware()))
	http.HandleFunc("GET /users/{id}", middlewares.ChainMiddlewares(userHandler.Show, middlewares.AuthMiddleware()))
	http.HandleFunc("POST /users", middlewares.ChainMiddlewares(userHandler.Create, middlewares.AuthMiddleware()))
	http.HandleFunc("PUT /users/{id}", middlewares.ChainMiddlewares(userHandler.Update, middlewares.AuthMiddleware()))
	http.HandleFunc("DELETE /users/{id}", middlewares.ChainMiddlewares(userHandler.Delete, middlewares.AuthMiddleware()))

	http.HandleFunc("POST /auth/login", authHandler.Login)
	http.HandleFunc("POST /auth/register", authHandler.Register)

	if err := http.ListenAndServeTLS(":8080", "certs/localhost.pem", "certs/priv.key", nil); err != nil {
		panic(err)
	}
}
