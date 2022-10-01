package main

import "github.com/vklokov/keystore/core"

func main() {
	app := core.New()
	app.Start(3000)
}
