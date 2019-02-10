package main

import "github.com/koding/kite"

func main() {
	k := kite.New("user", "1.0.0")
	k.Config.Port = 6666
	k.Run()
}