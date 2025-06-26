package cmd

import (
	"context"
	"fmt"
	"github.com/fataevalex/k8s-controller/pkg/ctrl"
	"github.com/fataevalex/k8s-controller/pkg/informer"
	"github.com/go-logr/zerologr"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/valyala/fasthttp"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	ctrlruntime "sigs.k8s.io/controller-runtime"
	controllerlog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
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

		// Start controller-runtime manager and controller
		controllerlog.SetLogger(zerologr.New(&log.Logger))
		mgr, err := ctrlruntime.NewManager(ctrlruntime.GetConfigOrDie(), manager.Options{})
		if err != nil {
			log.Error().Err(err).Msg("Failed to create controller-runtime manager")
			os.Exit(1)
		}
		if err := ctrl.AddDeploymentController(mgr); err != nil {
			log.Error().Err(err).Msg("Failed to add deployment controller")
			os.Exit(1)
		}
		go func() {
			log.Info().Msg("Starting controller-runtime manager...")
			if err := mgr.Start(cmd.Context()); err != nil {
				log.Error().Err(err).Msg("Manager exited with error")
				os.Exit(1)
			}
		}()

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
			requestID := uuid.New().String()
			ctx.Response.Header.Set("X-Request-ID", requestID)
			logger := log.With().Str("request_id", requestID).Logger()
			switch string(ctx.Path()) {
			case "/deployments":
				logger.Info().Msg("Deployments request received")
				ctx.Response.Header.Set("Content-Type", "application/json")
				deployments := informer.GetDeploymentNames()
				logger.Info().Msgf("Deployments: %v", deployments)
				ctx.SetStatusCode(200)
				ctx.Write([]byte("["))
				for i, name := range deployments {
					ctx.WriteString("\"")
					ctx.WriteString(name)
					ctx.WriteString("\"")
					if i < len(deployments)-1 {
						ctx.WriteString(",")
					}
				}
				ctx.Write([]byte("]"))
				return
			default:
				logger.Info().Msg("Default request received")
				fmt.Fprintf(ctx, "Hello from FastHTTP!")
			}
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
