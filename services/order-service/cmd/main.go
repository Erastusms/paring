package main

import (
	"log"

	"order-service/internal/handler"
	"order-service/internal/middlewares"
	"order-service/internal/repository"
	"order-service/internal/service"
	"order-service/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	gormDB := db.ConnectDB()
	repo := repository.NewOrderRepository(gormDB)
	svc := service.NewOrderService(repo)
	h := handler.NewOrderHandler(svc)

	r := gin.Default()
	r.Use(middlewares.JSONResponseMiddleware())
	r.Use(middlewares.JSONRecovery())
	
	api := r.Group("/api/orders")
	api.Use(h.AuthMiddleware)  // Apply auth to all /api/orders/*
	{
		api.POST("", h.CreateOrderHandler)
		api.GET("", h.GetOrdersHandler)
		api.GET("/:id", h.GetOrderHandler)
	}

	log.Println("Order Service running on :8082")
	if err := r.Run(":8082"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}