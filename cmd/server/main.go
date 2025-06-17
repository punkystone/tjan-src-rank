package main

import (
	"tjan-src-rank/internal/server"
	"tjan-src-rank/internal/src"
)

func main() {
	srcAPI := src.New()
	err := server.StartServer(srcAPI)
	if err != nil {
		panic(err)
	}
}
