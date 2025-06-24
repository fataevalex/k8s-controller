package cmd

import (
	"context"
	"fmt"
	"github.com/fataevalex/k8s-controller/pkg/informer"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

var serverPort int
var serverInCluster bool

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a FastHTTP server and deployment informer",
	Run: func(cmd *cobra.Command, args []string) {

		clientset, err := getServerKubeClient(KubeConfigPath, serverInCluster)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create Kubernetes client")
			os.Exit(1)
		}

		ctx := context.Background()
		go informer.StartDeploymentInformer(ctx, clientset)

		handler := func(ctx *fasthttp.RequestCtx) {

			requestLog := log.Debug()
			requestLog = requestLog.
				Str("method", string(ctx.Method())).
				Str("request_uri", string(ctx.RequestURI())).
				Str("path", string(ctx.Path())).
				Str("host", string(ctx.Host())).
				Str("remote_addr", ctx.RemoteAddr().String()).
				Str("local_addr", ctx.LocalAddr().String()).
				Str("protocol", string(ctx.Request.Header.Protocol())).
				Bool("is_tls", ctx.IsTLS()).
				Str("user_agent", string(ctx.UserAgent()))
			queryParams := make(map[string]string)
			ctx.QueryArgs().VisitAll(func(key, value []byte) {
				queryParams[string(key)] = string(value)
			})
			if len(queryParams) > 0 {
				requestLog = requestLog.Interface("query_params", queryParams)
			}
			requestHeaders := make(map[string]string)
			ctx.Request.Header.VisitAll(func(key, value []byte) {
				requestHeaders[string(key)] = string(value)
			})
			if len(requestHeaders) > 0 {
				requestLog = requestLog.Interface("request_headers", requestHeaders)
			}
			requestLog.Send()
			fmt.Fprintf(ctx, "Hello from FastHTTP!")
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
	InitConfig()
	serverCmd.Flags().IntVar(&serverPort, "port", 8080, "Port to run the server on")
	serverCmd.Flags().BoolVar(&serverInCluster, "in-cluster", false, "Use in-cluster Kubernetes config")
}

func getServerKubeClient(kubeconfigPath string, inCluster bool) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	if inCluster {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	}
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
