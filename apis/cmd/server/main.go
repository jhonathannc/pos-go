package main

import "github.com/jhonathannc/pos-go/apis/configs"

func main() {
	cfg, _ := configs.LoadConfig(".")
	println(cfg.JwtExpiresIn)
}
