package main

import (
	"BriefRetelling2.0/config"
	"fmt"
)

func main() {
	cfg := config.MustLoadConfig()
	fmt.Printf("%#v\n", cfg)
	// TODO: bot

	// TODO: GPT

	// TODO: Run server
}
