// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()

// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the Poker App!"})
// 	})

// 	r.Run(":8080")
// }

package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Enable CORS for development
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	routes.setupRoutes(r)
	r.Run(":8080")
}
