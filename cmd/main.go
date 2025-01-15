package main

import (
	"log"

	"github.com/RaihanArdiann/BEproject_simpleForum/internal/configs"
	"github.com/RaihanArdiann/BEproject_simpleForum/internal/handlers/memberships"
	membershipRepo "github.com/RaihanArdiann/BEproject_simpleForum/internal/repository/memberships"
	"github.com/RaihanArdiann/BEproject_simpleForum/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("failed to initialize config")
	}
	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	_ = membershipRepo.NewRepository(db)

	membershipsHandler := memberships.NewHandler(r)
	membershipsHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
