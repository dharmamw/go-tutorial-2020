package main

import (
	"fmt"

	"tugas-arif/internal/boot"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := boot.HTTP(); err != nil {
		fmt.Println("[HTTP] failed to boot http server due to " + err.Error())
	}
}
