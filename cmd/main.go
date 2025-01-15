package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan.ardiann/simpleForum/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()

	r.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
