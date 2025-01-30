package main

import (
	"log"

	"github.com/RaihanArdiann/BEproject_simpleForum/internal/configs"
	"github.com/RaihanArdiann/BEproject_simpleForum/internal/handlers/memberships"
	"github.com/RaihanArdiann/BEproject_simpleForum/internal/handlers/posts"
	membershipRepo "github.com/RaihanArdiann/BEproject_simpleForum/internal/repository/memberships"
	postRepo "github.com/RaihanArdiann/BEproject_simpleForum/internal/repository/posts"
	membershipSvc "github.com/RaihanArdiann/BEproject_simpleForum/internal/services/memberships"
	postSvc "github.com/RaihanArdiann/BEproject_simpleForum/internal/services/posts"
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
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipsHandler := memberships.NewHandler(r, membershipService)
	postHandler := posts.NewHandler(r, postService)

	membershipsHandler.RegisterRoute()
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
