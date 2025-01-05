package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }
        c.Next()
    }
}

func Handler(c *gin.Context) {
    gin.SetMode(gin.ReleaseMode)
    app := config.App()
    env := app.Env
    gin := gin.Default()
    gin.Use(CORSMiddleware())
    db := app.Db
    route.Setup(env, db, gin)
    gin.ServeHTTP(c.Writer, c.Request)
}