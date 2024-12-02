package main

import (
	"context"
	"fmt"
	"github.com/diyorich/post-api/internal/config"
	"github.com/diyorich/post-api/internal/handler"
	"github.com/diyorich/post-api/internal/repository/cache"
	repo "github.com/diyorich/post-api/internal/repository/post"
	"github.com/diyorich/post-api/internal/service/post"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	ctx := context.Background()

	redis, err := cache.Dial(ctx, cfg.Cache)
	if err != nil {
		return err
	}
	repository := repo.NewRepository(redis)
	service := post.NewService(repository)

	fmt.Println("Preparing posts data...")
	ctx, _ = context.WithTimeout(ctx, time.Second*10)
	if err := service.Load(ctx, "./data/MOCK_DATA.json"); err != nil {
		return errors.Wrap(err, "error on preparing data")
	}
	fmt.Println("Posts data loaded successfully")

	h := handler.NewHandler(service)
	r := gin.Default()
	r.GET("/external/api/posts", h.GetList)

	err = r.Run(fmt.Sprintf(":%v", cfg.App.Port))
	if err != nil {
		return err
	}

	return nil
}
