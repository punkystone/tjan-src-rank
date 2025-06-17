package main

import (
	"tjan-src-rank/internal/server"
	"tjan-src-rank/internal/src"
	"tjan-src-rank/internal/util"
)

func main() {
	env, err := util.CheckEnv()
	if err != nil {
		panic(err)
	}
	srcAPI, err := src.New(env.User)
	if err != nil {
		panic(err)
	}
	err = server.StartServer(srcAPI)
	if err != nil {
		panic(err)
	}
}
