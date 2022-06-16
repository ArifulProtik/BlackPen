package main

import (
	"log"
	"time"

	"github.com/ArifulProtik/BlackPen/config"
	"github.com/ArifulProtik/BlackPen/ent"
	"github.com/ArifulProtik/BlackPen/pkg/auth"
	"github.com/ArifulProtik/BlackPen/pkg/router"
	"github.com/ArifulProtik/BlackPen/pkg/server"
	"github.com/ArifulProtik/BlackPen/pkg/services"
)

func main() {
	cfg, err := config.New("./", "app.env", "env")
	if err != nil {
		log.Fatal(err)
	}

	server := server.New(&cfg.ServerConfig)
	newdb := ent.NewdbClient(&cfg.Postegres)
	mainservice := services.New(newdb)
	exp := time.Minute * 10
	Auth := auth.NewToken(cfg.ServerConfig.Jwt_Key, exp)
	log.Println(cfg.Jwt_Key)
	router.InitRouter(server.Echo.Group("/api/v1"), mainservice, Auth)
	server.Run()
}
