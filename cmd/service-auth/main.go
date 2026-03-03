package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/auth/login", handleLogin)
	r.POST("/auth/register", handleRegister)
	r.POST("/auth/refresh", handleRefresh)
	r.POST("/auth/logout", handleLogout)
	r.GET("/auth/me", handleMe)
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "healthy"}) })
	http.ListenAndServe(":8082", r)
}

func handleLogin(c *gin.Context) { c.JSON(200, gin.H{"token": "jwt-token"}) }
func handleRegister(c *gin.Context) { c.JSON(201, gin.H{"message": "registered"}) }
func handleRefresh(c *gin.Context) { c.JSON(200, gin.H{"token": "new-token"}) }
func handleLogout(c *gin.Context) { c.JSON(200, gin.H{"message": "logged out"}) }
func handleMe(c *gin.Context) { c.JSON(200, gin.H{"id": "user-1", "email": "user@flowforge.io"}) }
