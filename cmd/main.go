package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/raihan.ardiann/simpleForum/internal/configs"
	"github.com/raihan.ardiann/simpleForum/internal/handlers/memberships"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("failed to initialize config")
	}
	cfg = configs.Get()
	log.Println("config", cfg)

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()

	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
