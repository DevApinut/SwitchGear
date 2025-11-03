package main

import (
	"SwitchGear/internal/keycloak/auth"
	"SwitchGear/internal/keycloak/config"
	authhandler "SwitchGear/internal/keycloak/handler/auth"
	"SwitchGear/internal/keycloak/handler/render"
	"SwitchGear/internal/keycloak/middleware"
	rds "SwitchGear/internal/keycloak/store/redis"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatalf("failed to load and parse config : %v", err)
		return
	}
	serverAddr := fmt.Sprintf("%s:%s", cfg.AppHost, cfg.AppPort)
	// context
	ctx := context.Background()

	authOptions := []auth.Option{
		auth.WithClientSecret(cfg.Auth.ClientSecret),
		auth.WithRealmKeycloak(cfg.Auth.Realm),
	}

	authClient, err := auth.New(
		ctx,
		cfg.Auth.BaseURL,
		cfg.Auth.ClientID,
		cfg.Auth.RedirectURL,
		authOptions...,
	)
	if err != nil {
		log.Fatalf("Failed to initialize auth client : %v", err)
		return
	}

	// Redis
	redisClient := redis.NewClient(&cfg.RedisConfig)
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
		return
	}
	defer redisClient.Close()
	router := gin.Default()
	// Load HTML templates from internal/templates
	// Using relative path from where you run the application
	router.LoadHTMLGlob("./internal/keycloak/templates/*/*.tmpl")

	authStore := rds.NewAuthRedisManager(redisClient)
	sessionStore := rds.NewSessionRedisManager(redisClient)
	authHandler := authhandler.New(cfg,
		serverAddr,
		authClient,
		authStore,
		sessionStore,
	)
	renderHandler := render.New(cfg)

	group := router.Group("/api/v1")
	group.GET("/login", authHandler.RenderLoginPage)
	group.GET("/login-keycloak", authHandler.RedirectToKeycloak)
	// r.GET("/login-keycloak", authHandler.RedirectToKeycloak)
	group.GET("/logout", authHandler.Logout)
	group.GET("/callback-logout", authHandler.CallbackLogout)
	group.GET("/callback", authHandler.Callback)
	// r.GET("/callback-auth", authHandler.Callback)

	// Protected routes
	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware(sessionStore, authClient))
	{
		protected.GET("/success-login", renderHandler.SuccessLogin)
		// Add other protected routes here
	}
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
