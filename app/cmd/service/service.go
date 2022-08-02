package main

import (
	"fmt"
	"github.com/eznd-otus-msa/hw3/app/internal/transport/server/httpmw"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jinzhu/configor"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dblogger "gorm.io/gorm/logger"

	"github.com/eznd-otus-msa/hw3/app/internal/domain"
	"github.com/eznd-otus-msa/hw3/app/internal/service"
	"github.com/eznd-otus-msa/hw3/app/internal/transport/client/dbrepo"
	"github.com/eznd-otus-msa/hw3/app/internal/transport/server/http"
	_ "github.com/eznd-otus-msa/hw3/app/migrations"
)

func main() {
	var cfg domain.Config
	err := configor.New(&configor.Config{Silent: true}).Load(&cfg, "config/config.yaml", "./config.yaml")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbrepo.Dsn(cfg.Db),
	}), &gorm.Config{
		Logger: dblogger.Default.LogMode(dblogger.Info),
	})
	if err != nil {
		panic(err)
	}

	userRepo := dbrepo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	getUserHandler := http.NewGetUser(userService)
	postUserHandler := http.NewPostUser(userService)
	putUserHandler := http.NewPutUser(userService)
	patchUserHandler := http.NewPatchUser(userService)
	deleteUserHandler := http.NewDeleteUser(userService)

	srv := fiber.New(fiber.Config{})

	prometheus := httpmw.New("otus-msa-hw3")
	prometheus.RegisterAt(srv, "/metrics")
	srv.Use(prometheus.Middleware)

	srv.Use(logger.New())
	srv.Use(favicon.New())
	srv.Use(recover.New())
	srv.Use(httpmw.NewChaosMonkeyMw())

	api := srv.Group("/api")

	api.Post("/user", postUserHandler.Handle())
	api.Post("/users", postUserHandler.Handle())

	api.Get("/user/:id", getUserHandler.Handle())
	api.Get("/users/:id", getUserHandler.Handle())

	api.Put("/user/:id", putUserHandler.Handle())
	api.Put("/users/:id", putUserHandler.Handle())

	api.Patch("/user/:id", patchUserHandler.Handle())
	api.Patch("/users/:id", patchUserHandler.Handle())

	api.Delete("/user/:id", deleteUserHandler.Handle())
	api.Delete("/users/:id", deleteUserHandler.Handle())

	srv.Get("/probe/live", http.RespondOk)
	srv.Get("/probe/ready", http.RespondOk)

	srv.All("/*", http.DefaultResponse)

	err = srv.Listen(fmt.Sprintf(":%s", cfg.Http.Port))
	if err != nil {
		panic(err)
	}
}
