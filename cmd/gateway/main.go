package main

import "github.com/cold-runner/Hylark/internal/gateway/gw"

func main() {
	g := gw.New()
	g.Spin()
}
