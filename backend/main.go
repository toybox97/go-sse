package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/toybox97/go-sse/config"
	"github.com/toybox97/go-sse/routes"
)

func main() {
	cfg := config.LoadConfig()
	
	r := gin.Default()
	
	routes.SetupRoutes(r)
	
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(cfg.GetAddr()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
