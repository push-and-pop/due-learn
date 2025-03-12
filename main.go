package main

import (
	"due-learn/server/game"
	"due-learn/server/gate"
	"due-learn/server/login"
	"flag"
)

var serverType = flag.String("type", "", "server type:gate game")

func main() {
	flag.Parse()

	switch *serverType {
	case "gate":
		gate.Init()
	case "login":
		login.Init()
	case "game":
		game.Init()

	default:
		panic("err server type")
	}
}
