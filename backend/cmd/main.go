package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	mwsclient "d/internal/adapters/http-clients/mws-client"
	"d/internal/adapters/postgresql"
	"d/internal/config"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg := config.MustLoadConfig()

	postgres, err := postgresql.ConnectPostgreSQL(cfg.PostgesConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer postgres.Close()

	postgrestrg := postgresql.NewStorage(postgres, logger)

	mws := mwsclient.NewMWSClient(logger, cfg.HTTPClientsConfig)

	res, err := mws.TakeAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	if err := postgrestrg.AddMessage(postgresql.Message{ChatNum: 2, IsUser: true, Message: "Привет, asdfasdffasddasfу меня вопрос."}); err != nil {
		fmt.Println("first add", err)
		return
	}
	if err := postgrestrg.AddMessage(postgresql.Message{ChatNum: 2, IsUser: false, Message: "Задаfasdfsdaвай"}); err != nil {
		fmt.Println("sec add", err)
		return
	}
	if err := postgrestrg.AddMessage(postgresql.Message{ChatNum: 2, IsUser: true, Message: "Последнее fsdafsdaвидео какое там популярное"}); err != nil {
		fmt.Println("th add", err)
		return
	}
	if err := postgrestrg.AddMessage(postgresql.Message{ChatNum: 2, IsUser: false, Message: "это: ,..fdsadfsafd.."}); err != nil {
		fmt.Println("fo add", err)
		return
	}
	if err := postgrestrg.AddMessage(postgresql.Message{ChatNum: 2, IsUser: true, Message: "Спасибо, еще вернусьfdfadsfd"}); err != nil {
		fmt.Println("fi add", err)
		return
	}

	mes, err := postgrestrg.TakeByChatNum(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("вывод смс")
	for _, m := range mes {
		fmt.Println(m)
	}
}
