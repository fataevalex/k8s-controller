/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Kubernetes deployments in the namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("k8s-controller delete [deployment-name] command executed.")
		InitConfig()
		DeploymentName := args[0]
		log.Debug().Msgf("Using namespace %s", Namespace)
		log.Debug().Msgf("Using config %s", KubeConfigPath)
		log.Debug().Msgf("Deploment to delete %s", DeploymentName)

		if AllNamespaces {
			log.Error().Msg("Deleting across all namespaces is not supported. Please specify --namespace.")
			os.Exit(1)
		}

		clientset, err := getKubeClient(KubeConfigPath)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create Kubernetes client")
			os.Exit(1)
		}

		err = clientset.AppsV1().Deployments(Namespace).Delete(context.Background(), DeploymentName, metav1.DeleteOptions{})
		if err != nil {
			log.Error().Err(err).Msgf("Failed to delete deployment %s", DeploymentName)
			os.Exit(1)
		}
		fmt.Fprintf(cmd.OutOrStdout(), "Deployment '%s' deleted from namespace '%s'\n", DeploymentName, Namespace)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
