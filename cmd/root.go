/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time" // Add import time to time formater

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log" // log.Logger - это глобальный логгер zerolog
	"github.com/spf13/cobra"
)

var logLevel string // global variable to store --log-level

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8s-controller",
	Short: "k8s-controller is simple kubernetes controller on golang",
	Long: `
k8s-controller is a CLI tool to manage Kubernetes resources.
It allows you to perform various operations on your cluster.`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("k8s-controller root command executed.")
		// Base logic
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Init cobra to parsing input params
	cobra.OnInitialize(initLogger)

	// Set  --log-level. flag to all application
	rootCmd.PersistentFlags().StringVar(
		&logLevel,   // global loglevel
		"log-level", // loglevel  param name
		"info",      // set default log level
		fmt.Sprintf("Set log level: %s", getValidZerologLevels()),
	)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", Namespace, "Kubernetes namespace")
	rootCmd.PersistentFlags().StringVar(&KubeConfigPath, "kubeconfig", "", "Path to kubeconfig file")
	rootCmd.PersistentFlags().BoolVarP(&AllNamespaces, "all-namespaces", "A", false, "If present, list across all namespaces")
}

// initLogger will be executed by Cobra after parsing command line params.
func initLogger() {
	parsedLevel := parseLogLevel(logLevel) // set loglevel form command line parmas
	configureLogger(parsedLevel)           // execute logger configurator

	log.Debug().Msgf("Logger initialized with level: %s", logLevel) //
}

// parseLogLevel translate command line loglevel  в zerolog.Level format
func parseLogLevel(lvl string) zerolog.Level {
	switch strings.ToLower(lvl) {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "disabled":
		return zerolog.Disabled
	default:
		fmt.Fprintf(os.Stderr, "Warning: Invalid log level '%s' provided. Using 'info' level. Valid levels are: %s\n", lvl, getValidZerologLevels())
		return zerolog.InfoLevel
	}
}

// configureLogger set global zerolog on setting log level
func configureLogger(level zerolog.Level) {
	// set global log level
	zerolog.SetGlobalLevel(level)
	// set time format
	zerolog.TimeFieldFormat = time.RFC3339Nano // use RFC3339Nano to microsecond

	//set logger to output to console
	//consoleWriter := zerolog.ConsoleWriter{
	//	Out:        os.Stderr,
	//	TimeFormat: "2006-01-02 15:04:05.000", // datetime format for console
	//}

	// fine tuning Caller
	//if level == zerolog.TraceLevel || level == zerolog.DebugLevel {
	//	// For Trace and Debug, turn on  Caller and setup order for Caller
	//	consoleWriter.PartsOrder = []string{
	//		zerolog.TimestampFieldName,
	//		zerolog.LevelFieldName,
	//		zerolog.CallerFieldName,
	//		zerolog.MessageFieldName,
	//	}
	//	// tuning formating Caller
	//	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
	//		// get file name and string number
	//		short := file
	//		for i := len(file) - 1; i > 0; i-- {
	//			if file[i] == '/' {
	//				short = file[i+1:]
	//				break
	//			}
	//		}
	//		file = short
	//		return fmt.Sprintf("%s:%d", file, line)
	//	}
	//	// setup global logger wich ConsoleWriter and Caller
	//	log.Logger = zerolog.New(consoleWriter).With().Timestamp().Caller().Logger().Level(level) //
	//} else {
	//	// To loglevel above Debug, Caller обычно не нужен для производительности и чистоты логов
	//	consoleWriter.PartsOrder = []string{
	//		zerolog.TimestampFieldName,
	//		zerolog.LevelFieldName,
	//		zerolog.MessageFieldName,
	//	}
	//	// setup global logger ConsoleWriter without Caller
	//	log.Logger = zerolog.New(consoleWriter).With().Timestamp().Logger().Level(level) //
	//}

	// redirect standart logger zerolog (log.Logger)
}

// getValidZerologLevels - helper function
func getValidZerologLevels() string {
	levels := []string{
		zerolog.TraceLevel.String(),
		zerolog.DebugLevel.String(),
		zerolog.InfoLevel.String(),
		zerolog.WarnLevel.String(),
		zerolog.ErrorLevel.String(),
		zerolog.FatalLevel.String(),
		zerolog.PanicLevel.String(),
		zerolog.Disabled.String(),
	}
	return strings.Join(levels, ", ")
}
