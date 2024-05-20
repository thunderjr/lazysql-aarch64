package main

import (
	"io"
	"log"

	"github.com/thunderjr/lazysql-aarch64/app"
	"github.com/thunderjr/lazysql-aarch64/components"

	"github.com/go-sql-driver/mysql"
)

func main() {
	mysql.SetLogger(log.New(io.Discard, "", 0))

	if err := app.App.
		SetRoot(components.MainPages, true).
		EnableMouse(true).
		Run(); err != nil {
		panic(err)
	}
}
