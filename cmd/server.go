package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
)

var serverPort int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a FastHTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		handler := func(ctx *fasthttp.RequestCtx) {
			hostname := string(ctx.Host())
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02 15:04:05 MST")

			fmt.Fprintf(ctx, "Hello from FastHTTP!\n")
			fmt.Fprintf(ctx, "Hostname: %s\n", hostname)
			fmt.Fprintf(ctx, "Current Time: %s\n", formattedTime)
		}
		addr := fmt.Sprintf(":%d", serverPort)
		log.Info().Msgf("Starting FastHTTP server on %s", addr)
		if err := fasthttp.ListenAndServe(addr, handler); err != nil {
			log.Error().Err(err).Msg("Error starting FastHTTP server")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVar(&serverPort, "port", 8080, "Port to run the server on")
}
