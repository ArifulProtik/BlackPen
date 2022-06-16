package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ArifulProtik/BlackPen/ent/migrate"

	cfg "github.com/ArifulProtik/BlackPen/config"
	_ "github.com/lib/pq"
)

func NewdbClient(c *cfg.Postegres) *Client {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_NAME, c.DB_PASS)

	client, err := Open("postgres", psqlInfo)
	if err != nil {
		log.Panicln(err)
	}
	log.Println("DB Connected")
	if err := client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true)); !errors.Is(err, nil) {
		log.Println(err)
	}
	return client
}
