/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
// "github.com/fataevalex/k8s-controller/cmd"
"github.com/rs/zerolog/log"
)

func main() {
 log.Info().Msg("info message")
    log.Debug().Msg("debug message")
    log.Trace().Msg("trace message")
    log.Warn().Msg("warn message")
    log.Error().Msg("error message")
// 	cmd.Execute()
}
