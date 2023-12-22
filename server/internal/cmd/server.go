package main

import "github.com/gin-gonic/gin"

type Server struct {
	port   string
	Router *gin.Engine
}
