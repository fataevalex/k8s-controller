/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/fataevalex/k8s-controller/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Debug().Msg("Starting main")
	cmd.Execute()
}
