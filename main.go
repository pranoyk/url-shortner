package main

import "main/router"

func main() {
	startService()
}

func startService() {
	engine := router.Init()
	engine.Run()
}