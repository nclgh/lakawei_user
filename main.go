package main

import "github.com/koding/kite"

func main() {
	k := kite.New("kite", "1.0.0")
	k.Config.Port = 8080
	k.Run()
}