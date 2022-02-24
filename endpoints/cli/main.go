package main

import (
	"github.com/joho/godotenv"
	"github.com/yaroslavklimuk/crazy-lottery/cli_handler"
	"github.com/yaroslavklimuk/crazy-lottery/storage"
	"os"
)

func main() {
	if os.Args[1] == "" {
		panic("Pass an action {process_money|process_items} as an argument")
	}
	action := os.Args[1]

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbFile := os.Getenv("DB_FILE")
	if dbFile == "" {
		panic("Bad database configuration")
	}
	dataStorage, err := storage.GetStorage(dbFile)
	if err != nil {
		panic(err)
	}

	handler := cli_handler.MakeCliHandler(dataStorage)

	switch action {
	case cli_handler.ActionProcessMoney:
		err = handler.ProcessMoneyRewards()
		if err != nil {
			panic(err)
		}
	case cli_handler.ActionProcessItems:
		err = handler.ProcessItemRewards()
		if err != nil {
			panic(err)
		}
	default:
		panic("Pass an action {process_money|process_items} as an argument")
	}
}
